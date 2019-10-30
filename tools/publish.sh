#!/bin/bash

export GIT_MERGE_AUTOEDIT=no

VERSION=$1

set -e

git flow release start ${VERSION}
git flow release finish -m "${VERSION}" ${VERSION}
git push --all
git checkout develop
