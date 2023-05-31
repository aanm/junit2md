FROM docker.io/library/golang:1.20.4@sha256:690e4135bf2a4571a572bfd5ddfa806b1cb9c3dea0446ebadaf32bc2ea09d4f9 as builder
WORKDIR /src/junit2md
COPY . .
RUN make junit2md

FROM docker.io/library/busybox:1.36.1@sha256:560af6915bfc8d7630e50e212e08242d37b63bd5c1ccf9bd4acccf116e262d5b
COPY --from=builder /src/junit2md/junit2md /usr/bin/junit2md
