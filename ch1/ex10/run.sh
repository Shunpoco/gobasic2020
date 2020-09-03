#!/bin/bash

: Tune shell options && {
  set -o errexit
  set -o nounset
  set -o xtrace
}

: Set Variables && {
  readonly SITES="https://www.youtube.com/"
}
: Run go code && {
  go run main.go ${SITES}
}
