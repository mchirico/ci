


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

### Must git commit files..

Concourse will pull the *.yml* files from the repo, so these
files must be checked to your repo on github.  




By default, this will pull from the checked in files from the *master* repo. If you
take a look at the first few lines of *build-golang-pipeline.yml*, you'll see
*branch: master* listed.  You can change this, if you want to run Concourse
on a different branch.

```bash
$ tail build-golang-pipeline.yml

resources:

- name: date
  type: git
  source:
    uri: https://github.com/mchirico/date.git
    branch: master


```


### Once checked in

Now, if you run *ci/run_ci.sh*, on your local branch. 

```bash
cd ci
./run_ci.sh
```

You should see the pipeline in Concourse created.




<a href="https://mchirico.github.io/p/images/ciEx1.png">
<img src="https://mchirico.github.io/p/images/ciEx1.png" width="650"/>
</a>
