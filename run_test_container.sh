repo="github.com/waieez/fbforbots"
docker run -v "$PWD:/go/src/$repo" --rm -w "/go/src/$repo" golang:1.6.2 /bin/bash -c "./container_test.sh"