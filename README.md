


[![Build Status](https://travis-ci.org/mchirico/ci.svg?branch=master)](https://travis-ci.org/mchirico/ci)
[![codecov](https://codecov.io/gh/mchirico/ci/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/ci)
# ci

ci is a program for building concourse pipeline files for your go project.  

Here's a complete, simple example. Assume project [**github.com/mchirico/date.git**](https://github.com/mchirico/date) doesn't have a *ci* directory
with the concourse yml files.

```bash

git clone https://github.com/mchirico/date.git
cd date

# The following command will create a directory ci
ci create
```

Now the directory *ci* has the following files:

```bash
$ ls ci
build-golang-pipeline.yml	inform-task.yml
build-task.yml			inform.sh
build.sh			run_ci.sh
docker-task.yml			unit-task.yml
docker.sh			unit.sh
```

Now, if you run run_ci.sh, you'll start the following pipeline:

```bash
cd ci
./run_ci.sh
```



<a href="https://mchirico.github.io/p/images/ciEx1.png">
<img src="https://mchirico.github.io/p/images/ciEx1.png" width="650"/>
</a>
