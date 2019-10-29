#!/bin/bash

rm -rf dist
gox -output="dist/codeon_{{.OS}}_{{.Arch}}"
