#!/bin/bash -ex

echo "==> Getting code coverage tools"
go get -v github.com/jstemmer/go-junit-report

echo "==> Getting dependencies"
go get -v ./...

echo "==> Running tests and gathering coverage"
rm -f acc.out
for Dir in $(go list ./... | grep -v "/vendor/" );
do
    go test -v $Dir | tee -a acc.out
done

echo "==> Outputting coverage report"
cat acc.out | go-junit-report > coverage.xml

## Credit to: https://github.com/KevinPike