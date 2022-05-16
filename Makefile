DATE=$(shell date)
GOVERSION=$(shell go version | awk '{print $$3}')
BUILDVERSION=$(shell git describe --tags | awk '{print $$1}')
TAG=$(shell cat .version)
GOARCH= \
				amd64 \
				arm64

build:  linux darwin tar sha

linux: linux-amd64 linux-arm64

darwin: darwin-amd64 darwin-arm64

sha: sha-amd64 sha-arm64

tar: tar-amd64 tar-arm64

linux-%:
	go build -ldflags="-X 'main.BuildVersion=${BUILDVERSION}' -X 'main.BuildDate=${DATE}' -X 'main.Platform=linux/$*' -X 'main.GoVersion=${GOVERSION}'" -o kubectl_kopy_linux_$*/kubectl-kopy

darwin-%:
	go build -ldflags="-X 'main.BuildVersion=${BUILDVERSION}' -X 'main.BuildDate=${DATE}' -X 'main.Platform=darwin/$*' -X 'main.GoVersion=${GOVERSION}'" -o kubectl_kopy_darwin_$*/kubectl-kopy

sha-%:
	echo $(shell openssl sha256 < kubectl_kopy_linux_$*.tar.gz | awk '{print $$2}')	kubectl_kopy_linux_$* > kubectl_kopy_linux_$*.sha256
	echo $(shell openssl sha256 < kubectl_kopy_darwin_$*.tar.gz | awk '{print $$2}')	kubectl_kopy_darwin_$* > kubectl_kopy_darwin_$*.sha256

tar-%:
	cp LICENSE kubectl_kopy_linux_$* && cp LICENSE kubectl_kopy_darwin_$*
	cd kubectl_kopy_linux_$* && tar czf kubectl_kopy_linux_$*.tar.gz kubectl-kopy LICENSE && mv kubectl_kopy_linux_$*.tar.gz ../
	cd kubectl_kopy_darwin_$* && tar czf kubectl_kopy_darwin_$*.tar.gz kubectl-kopy LICENSE && mv kubectl_kopy_darwin_$*.tar.gz ../

clean:
	@rm -rf kubectl_kopy*

release:
	gh release create ${TAG} kubectl_kopy*.tar.gz kubectl_kopy*sha256 -F CHANGELOG/CHANGELOG-${TAG}.md

tag:
	git tag -a ${TAG} -m "version ${TAG}"
	git push origin ${TAG}
