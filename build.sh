#!/usr/bin/env bash

export PROVIDER_VERSION="0.0.2"

build() {
  export GOOS=$1
  export GOARCH=$2
  go build -o terraform-provider-dfd_v${PROVIDER_VERSION}
  chmod +x terraform-provider-dfd_v${PROVIDER_VERSION}
  tar -czvf terraform-provider-dfd_v${PROVIDER_VERSION}_${GOOS}-${GOARCH}.tar.gz terraform-provider-dfd_v${PROVIDER_VERSION}
}


build 'linux' 'amd64'
build 'linux' '386'

build 'windows' 'amd64'
build 'windows' '386'

build 'darwin' 'amd64'
