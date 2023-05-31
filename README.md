##  junit2md

`junit2md` is a program created to present the junit output files into a 
Markdown table.

This repository can also be used as GitHub action.  

## Usage

### Arguments

```bash
Usage of ./junit2md:
  -e    Show errored tests (default true)
  -f    Show failed tests (default true)
  -i string
        Input file path
  -o string
        Output file path
  -p    Show passed tests (default true)
  -s    Show skipped tests (default true)
  -v    Show stderr/stdout of tests
```

### Example

#### Successful case

```bash
$ ./junit2md -i example/successful/junit.xml -s=false -o example/successful/table.md
Markdown table saved to example/successful/table.md
```

#### With failures

```bash
$ ./junit2md -i example/failure/junit.xml -v=true -s=false -p=false -o example/failure/table.md
Markdown table saved to example/failure/table.md
```

## GitHub action usage

```yaml
name: Github Action Example

on:
  pull_request: {}

  push:
    branches:
      - main

jobs:
  test-junit-2-md:
    runs-on: ubuntu-latest
    name: Print JUnits as GitHub Summary
    steps:
      - name: Checkout code
        uses: actions/checkout@8e5e7e5ab8b370d6c329ec480221332ada57f0ab # v3.5.2
        with:
          ref: ${{ steps.vars.outputs.sha }}
          persist-credentials: false

      - name: Publish Test Results As GitHub Summary
        uses: aanm/junit2md@v0.0.1
        with:
          junit-directory: "example"

  test-junit-2-md-failing-tests:
    runs-on: ubuntu-latest
    name: Print JUnits Failures as GitHub Summary
    steps:
      - name: Checkout code
        uses: actions/checkout@8e5e7e5ab8b370d6c329ec480221332ada57f0ab # v3.5.2
        with:
          ref: ${{ steps.vars.outputs.sha }}
          persist-credentials: false

      - name: Publish Test Results As GitHub Summary
        uses: aanm/junit2md@v0.0.1
        with:
          junit-directory: "example"
```