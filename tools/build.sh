#!/bin/bash

rm -rf dist
gox -output="dist/codeopen_{{.OS}}_{{.Arch}}"
