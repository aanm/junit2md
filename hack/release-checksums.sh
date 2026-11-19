#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
#
# Builds the junit2md release binaries with the Go version in .go-version.

set -euo pipefail

usage() {
  cat >&2 <<EOF
usage: $0 [-c] [-w] [-o DIR]
  -c      fail if a sha256 differs from action.yaml
  -w      write the sha256 values into action.yaml
  -o DIR  output directory (default: dist)
EOF
  exit 2
}

check=false
write=false
out=
while getopts "cwo:" opt; do
  case "${opt}" in
    c) check=true ;;
    w) write=true check=true ;;
    o) out="${OPTARG}" ;;
    *) usage ;;
  esac
done
shift $((OPTIND - 1))
[ "$#" -eq 0 ] || usage

GO="${GO:-go}"
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
out="${out:-${root}/dist}"
mkdir -p "${out}"
out="$(cd "${out}" && pwd)"
cd "${root}"

# Keep the user's environment and go env file out of the build.
goenv() {
  env -i PATH="${PATH}" HOME="${HOME}" \
    GOENV=off GOFLAGS=-buildvcs=false GOPROXY=off GOTOOLCHAIN=local "$@"
}

want="go$(cat .go-version)"
have="$(goenv "${GO}" env GOVERSION)"
if [ "${have}" != "${want}" ]; then
  echo "$0: ${GO} is ${have}, .go-version pins ${want}; set GO to a ${want} binary" >&2
  exit 1
fi

for arch in amd64 arm64; do
  goenv CGO_ENABLED=0 GOOS=linux GOARCH="${arch}" GOAMD64=v1 GOARM64=v8.0 \
    "${GO}" build -trimpath -ldflags='-s -w' -o "${out}/junit2md-linux-${arch}" ./cmd
done
(cd "${out}" && sha256sum junit2md-linux-amd64 junit2md-linux-arm64 > SHA256SUMS)
cat "${out}/SHA256SUMS"

rc=0
while read -r sum asset; do
  arch="${asset#junit2md-linux-}"
  var="JUNIT2MD_SHA256_${arch^^}"
  if "${write}"; then
    sed -i "s/^\( *${var}: \).*/\1\"${sum}\"/" action.yaml
  fi
  pinned="$(sed -n "s/^ *${var}: *\"\(.*\)\"$/\1/p" action.yaml)"
  if "${check}" && [ "${pinned}" != "${sum}" ]; then
    echo "$0: ${asset} is ${sum}, action.yaml pins ${var}: ${pinned:-<missing>}" >&2
    rc=1
  fi
done < "${out}/SHA256SUMS"
exit "${rc}"
