#!/usr/bin/env bash 
set -xe
set GOOS=linux
set GOARCH=arm
set GOARM=5
# install packages and dependencies
#go tool dist install -v pkg/runtime
#go install -v -a std
#go get github.com/dgrijalva/jwt-go
# build command
#go build -o bin/application application.go
#!/usr/bin/env bash 
# build binary
GOARCH=amd64 GOOS=linux go build -o bin/application application.go

# create zip containing the bin, assets and .ebextensions folder
zip -r uploadThis.zip bin public .ebextensions