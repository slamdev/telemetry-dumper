#!/usr/bin/env bash

#
# Script fetches a template chart from bitnami https://github.com/bitnami/charts/tree/main/template/CHART_NAME
# Usage example: ./fetch.sh
#

set -o errexit
set -o nounset
set -o pipefail

# https://github.com/bitnami/charts/archive/refs/heads/main.zip

function fetch() {
  wget -q -O /tmp/charts-main.zip https://github.com/bitnami/charts/archive/refs/heads/main.zip
  unzip -qq -o /tmp/charts-main.zip -d /tmp
  cp -R /tmp/charts-main/template/CHART_NAME/ .
  rm -rf /tmp/charts-main*
  rm -rf README.md .helmignore
}

fetch "$@"
