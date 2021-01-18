ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
gen-mocks:
	cd ${ROOT_DIR}/internal/domains/repositories && mockery --all
test:
	go test -mod vendor  -p 1 `go list ./... | grep -v /vendor/ | grep  -v /utils/`
