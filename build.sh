#!/usr/bin/env bash
export MY_VERSION="$(echo $RANDOM | md5sum | head -c 5; echo)"
export MY_IMAGE="docker.sys-int.de/sys-int/cf-operator"
make generate
make manifests
make docker-build docker-push IMG="${MY_IMAGE}:${MY_VERSION}"
make deploy IMG="${MY_IMAGE}:${MY_VERSION}"
#make bundle IMG="${MY_IMAGE}:${MY_VERSION}"
#make bundle-build bundle-push BUNDLE_IMG="${MY_IMAGE}-bundle:${MY_VERSION}"
#operator-sdk cleanup cf-operator
#operator-sdk run bundle ${MY_IMAGE}-bundle:${MY_VERSION}