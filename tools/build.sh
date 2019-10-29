#!/bin/bash

rm -rf dist
gox -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
