#!/bin/bash

source rcs/github

export GIT_MERGE_AUTOEDIT=no

VERSION=$1

set -e

git flow release start ${VERSION}
git flow release finish -m "${VERSION}" ${VERSION}
git push --all

goreleaser

git checkout develop
