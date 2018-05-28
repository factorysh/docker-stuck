bin: vendor
	go build docker-stuck.go

pull:
	docker pull bearstech/golang-dep
	docker pull bearstech/upx

test: vendor
	go test github.com/factorysh/docker-stuck/containers

docker:
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

clean:
	rm -rf vendor
	rm -f docker-stuck
