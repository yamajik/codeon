#!/bin/bash

VERSION=`. tools/version.sh`

MAJOR=`echo ${VERSION} | cut -f 1 -d "."`
MINOR=`echo ${VERSION} | cut -f 2 -d "."`
MICRO=`echo ${VERSION} | cut -f 3 -d "."`

case $1 in
    major)
        ((MAJOR+=1))
        MINOR=0
        MICRO=0
        ;;
    minor)
        ((MINOR+=1))
        MICRO=0
        ;;
    *)
        ((MICRO+=1))
        ;;
esac

. tools/release.sh ${MAJOR}.${MINOR}.${MICRO}
