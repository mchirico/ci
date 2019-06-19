package templates

func UnitSH() string {
	const s = `#!/bin/bash

set -e -u -x

export GOPATH=$PWD/depspath:$PWD/gopath
export PATH=$PWD/depspath/bin:$PWD/gopath/bin:$PATH

cd gopath/src/{{.Path}}

echo
echo "Fetching dependencies..."
go get -v -t ./...

echo
echo "Running tests..."
go test -v -race ./...
`
	return s
}

func RunCI() string {
	const s = `#!/bin/bash

set -eu

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
export fly_target=${fly_target:-mce}
echo "Concourse API target ${fly_target}"
echo "Tutorial $(basename $DIR)"

pushd $DIR
  fly -t ${fly_target} set-pipeline -p {{.Reposhort}}-pipeline -c build-golang-pipeline.yml -n
  fly -t ${fly_target} unpause-pipeline -p {{.Reposhort}}-pipeline
#  fly -t ${fly_target} trigger-job -w -j tutorial-pipeline/job-hello-world
popd

echo -e "\n\n                  Common commands:"
echo -e "**************************************\n\n"
echo -e "\n"
echo -e "                           fly -t mce watch --job {{.Reposhort}}-pipeline/unit"
echo -e "                           fly -t mce builds|grep '{{.Reposhort}}-pipeline'"
echo -e "                           fly -t mce destroy-pipeline -p {{.Reposhort}}-pipeline -n"
echo -e "                           fly -t mce workers -d "
echo -e "                            "
echo -e "                           To login to a running container: "
echo -e "                           fly -t mce intercept --job {{.Reposhort}}-pipeline/unit "
echo -e "\n"
echo -e "\n"

`
	return s
}

func BuildUnit() string {
	const s = `platform: linux

image_resource:
  type: registry-image
  source: {repository: golang}

inputs:
- name: goscratch
  path: gopath/src/{{.Path}}

caches:
- path: depspath/
- path: gopath/pkg/

run:
  path: gopath/src/{{.Path}}/ci/unit.sh

`
	return s
}

func BuildSH() string {
	const s = `#!/bin/bash

set -e -u -x

export GOPATH=$PWD/depspath:$PWD/gopath
export PATH=$PWD/depspath/bin:$PWD/gopath/bin:$PATH

cd gopath/src/{{.Path}}

#cd cmd/cake

echo
echo "Fetching dependencies..."
go get -v -t ./...


echo
echo "Building..."
go build -v ./...

echo
echo "Smoke test..."
#./cake
`
	return s
}

func InformSH() string {
	const s = `#!/bin/bash

set -e -u -x

export GOPATH=$PWD/depspath:$PWD/gopath
export PATH=$PWD/depspath/bin:$PWD/gopath/bin:$PATH

cd gopath/src/{{.Path}}

#cd cmd/cake

echo
echo "Fetching dependencies..."
go get -v -t ./...


echo
echo "Building..."
go build -v ./...

echo
echo "Smoke test..."
#./cake
`
	return s
}

func DockerSH() string {
	const s = `#!/bin/bash

set -e -u -x

export GOPATH=$PWD/depspath:$PWD/gopath
export PATH=$PWD/depspath/bin:$PWD/gopath/bin:$PATH

cd gopath/src/{{.Path}}

#cd cmd/cake

echo
echo "Fetching dependencies..."
go get -v -t ./...


echo
echo "Building..."
go build -v ./...

echo
echo "Smoke test..."
#./cake
`
	return s
}

func UnitTaskYML() string {
	const s = `platform: linux

image_resource:
  type: registry-image
  source: {repository: golang}

inputs:
- name: {{.Reposhort}}
  path: gopath/src/{{.Path}}

caches:
- path: depspath/
- path: gopath/pkg/

run:
  path: gopath/src/{{.Path}}/ci/unit.sh
`
	return s
}

func BuildTask() string {
	const s = `
platform: linux

image_resource:
  type: registry-image
  source: {repository: golang}

inputs:
- name: {{.Reposhort}}
  path: gopath/src/{{.Path}}

caches:
- path: depspath/
- path: gopath/pkg/

run:
  path: gopath/src/{{.Path}}/ci/build.sh
`
	return s
}

func InformTask() string {
	const s = `
platform: linux

image_resource:
  type: registry-image
  source: {repository: golang}

inputs:
- name: {{.Reposhort}}
  path: gopath/src/{{.Path}}

caches:
- path: depspath/
- path: gopath/pkg/

run:
  path: gopath/src/{{.Path}}/ci/inform.sh
`
	return s
}

func DockerTask() string {
	const s = `
platform: linux

image_resource:
  type: registry-image
  source: {repository: golang}

inputs:
- name: {{.Reposhort}}
  path: gopath/src/{{.Path}}

caches:
- path: depspath/
- path: gopath/pkg/

run:
  path: gopath/src/{{.Path}}/ci/docker.sh
`
	return s
}

func Pipeline() string {
	const s = `
resources:

- name: {{.Reposhort}}
  type: git
  source:
    uri: {{.RepoHttp}}
    branch: {{.Branch}}

###############################################################################

jobs:

- name: unit
  plan:
  - get: {{.Reposhort}}
    trigger: true
  - task: unit
    file: {{.Reposhort}}/ci/unit-task.yml

- name: build
  plan:
  - get: {{.Reposhort}}
    trigger: true
    passed: [unit]
  - task: build
    file: {{.Reposhort}}/ci/build-task.yml

- name: inform
  plan:
  - get: {{.Reposhort}}
    trigger: true
    passed: [build]
  - task: inform
    file: {{.Reposhort}}/ci/inform-task.yml

- name: docker
  plan:
  - get: {{.Reposhort}}
    trigger: true
    passed: [inform]
  - task: docker
    file: {{.Reposhort}}/ci/docker-task.yml

`

	return s
}
