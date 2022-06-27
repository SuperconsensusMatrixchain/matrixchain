# init project PATH
HOMEDIR := $(shell pwd)
OUTDIR  := $(HOMEDIR)/output
COMPILECACHEDIR := $(HOMEDIR)/.compile_cache
XVMDIR  := $(COMPILECACHEDIR)/xvm
TESTNETDIR := $(HOMEDIR)/testnet

# init command params
export GO111MODULE=on
X_ROOT_PATH := $(HOMEDIR)
export X_ROOT_PATH
export PATH := $(OUTDIR)/bin:$(XVMDIR):$(PATH)

# docker 标签
DOCKER_TAG ?= latest
# docker user
DOCKER_USER ?= username
# docker password
DOCKER_PASSWORD ?= password


# make, make all
all: clean compile

# make compile, go build
compile: xvm xchain
xchain:
	bash $(HOMEDIR)/auto/build.sh

# make xvm
xvm:
	bash $(HOMEDIR)/auto/build_xvm.sh

# make test, test your code
test: xvm unit
unit:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

# make clean
cleanall: clean cleantest cleancache
clean:
	rm -rf $(OUTDIR)
cleantest:
	rm -rf $(TESTNETDIR)
cleancache:
	rm -rf $(COMPILECACHEDIR)


# deploy test network
testnet:
	bash $(HOMEDIR)/auto/deploy_testnet.sh

# avoid filename conflict and speed up build
.PHONY: all compile test clean


# copy docker scripts to OUTDIR
cp-docker.sh-to-outdir:
	cp ./deployment/docker_env_install.sh $(OUTDIR)


# make matrixchain for pushing dockerhub
matrixchain-docker-image-no-compile: cp-docker.sh-to-outdir
	sudo docker build -f ./matrixchain.Dockerfile -t superconsensuschain/matrixchain:${DOCKER_TAG} .

# login dockerhub
login-dockerhub:
	sudo docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD)

# push docker images to dockerhub
push-to-dockerhub: login-dockerhub
	sudo docker push superconsensuschain/matrixchain:${DOCKER_TAG}