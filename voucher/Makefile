BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client
PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])

SHELL := bash
SHELL_VERSION = $(shell echo $$BASH_VERSION)
UNAME := $(shell uname -s)
VERSION_AND_ARCH = $(shell uname -rm)
ifeq ($(UNAME),Darwin)
	OS = macos ${VERSION_AND_ARCH}
else ifeq ($(UNAME),Linux)
	OS = linux ${VERSION_AND_ARCH}
else
	$(error OS not supported by this Makefile)
endif
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
RM_F_CMD = rm -f
RM_RF_CMD = ${RM_F_CMD} -r
SERVER_BIN = ${SERVER_DIR}
CLIENT_BIN = ${CLIENT_DIR}



# profile: $@ 
# dog:
# 	echo ${PROTO_DIR}/${project}/proto/*.proto

# $(project):
# 	echo "------[Generating profile.proto...]-------"
# 	protoc -I${PROTO_DIR}/${project}/proto --go_out=. --go_opt=module=${PACKAGE} --go-grpc_out=. --go-grpc_opt=module=${PACKAGE}  ${PROTO_DIR}/${project}/proto/*.proto

pb:
	echo "------[Generating profile.proto...]-------"
	protoc -I${PROTO_DIR} --go_out=. --go_opt=module=${PACKAGE} --go-grpc_out=. --go-grpc_opt=module=${PACKAGE}  ${PROTO_DIR}/*.proto

# $(project):
# 	echo $@/${PROTO_DIR}/${project}/proto
# protoc grpc/profile.proto --go_opt=module=github.com/kaolnwza/muniverse/storage --go_out=. --go-grpc_opt=module=github.com/kaolnwza/muniverse/storage --go-grpc_out=.
# protoc grpc/profile.proto --go_out=. --go-grpc_out=.

# protoc -Igrpc/profile/proto --go_out=. --go_opt=module=github.com/kaolnwza/muniverse/storage --go-grpc_out=. --go-grpc_opt=module=github.com/kaolnwza/muniverse/storage grpc/profile/proto/profile.proto

migrate-up:
	echo "-----[Migrations UP...]-----"
	cd database && migrate -database $$DATABASE_URL -path migrations up