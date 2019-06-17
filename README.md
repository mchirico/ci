


[![Build Status](https://travis-ci.org/mchirico/ci.svg?branch=master)](https://travis-ci.org/mchirico/ci)
[![codecov](https://codecov.io/gh/mchirico/ci/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/ci)
# ci

## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```

