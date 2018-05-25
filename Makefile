bin: vendor
	go build docker-stuck.go

test: vendor
	go test github.com/factorysh/docker-stuck/containers

docker: vendor
	docker run -ti --rm \
	-v `pwd`:/go/src/github.com/factorysh/docker-stuck/ \
	-w /go/src/github.com/factorysh/docker-stuck/ \
    bearstech/golang-dep \
	make bin

upx:
	docker run -ti --rm \
	-v `pwd`:/upx \
	-w /upx \
	bearstech/upx \
	upx docker-stuck

vendor:
	dep ensure