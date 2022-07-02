# I usually keep a `VERSION` file in the root so that anyone
# can clearly check what's the VERSION of `master` or any
# branch at any time by checking the `VERSION` in that git
# revision.
#
# Another benefit is that we can pass this file to our Docker 
# build context and have the version set in the binary that ends
# up inside the Docker image too.
VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      cortiz/certview

clean:
	rm -rf ./dist

buildDeps:
	go install github.com/goreleaser/goreleaser@latest
goVendor:
	GOPROXY="" go mod vendor
build:
	goreleaser build --rm-dist --snapshot
vet:
	go vet ./...
tidy:
	GOPROXY="" go mod tidy
test:
	go test ./...

release:
	git tag -a -s $(VERSION) -m "Release" || true
	git push origin $(VERSION)

.PHONY: install test fmt release
