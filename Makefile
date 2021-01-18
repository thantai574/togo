ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
gen-mocks:
	cd ${ROOT_DIR}/internal/domains/repositories && mockery --all