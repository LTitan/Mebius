PROJECT=github.com/LTitan/Mebius
VERSION?=v1alpha1
PROJECT_APIS=${PROJECT}/apis/${VERSION}
PROTO_TYPES=${PROJECT}/pkg/protos/types
CLIENTSET=${PROJECT}/pkg/clients/clientset
INFORMER=${PROJECT}/pkg/clients/informer
LISTER=${PROJECT}/pkg/clients/lister

ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
	export GOPATH
endif
GOPATH_SRC=${GOPATH}/src

all: register-gen deepcopy-gen defaulter-gen openapi-gen client-gen lister-gen informer-gen

install-tools: goimports install-grpc-env
	go install k8s.io/code-generator/cmd/go-to-protobuf@v0.25.3
	go install k8s.io/code-generator/cmd/client-gen@v0.25.3
	go install k8s.io/code-generator/cmd/informer-gen@v0.25.3
	go install k8s.io/code-generator/cmd/deepcopy-gen@v0.25.3
	go install k8s.io/code-generator/cmd/lister-gen@v0.25.3
	go install k8s.io/code-generator/cmd/register-gen@v0.25.3
	go install k8s.io/code-generator/cmd/openapi-gen@v0.25.3
	go install k8s.io/code-generator/cmd/defaulter-gen@v0.25.3
	go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.10.0
	go install -mod=readonly github.com/gogo/protobuf/protoc-gen-gogo@v1.3.2
	go install -mod=readonly github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
	go install -mod=readonly github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
	go install -mod=readonly google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install -mod=readonly google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install -mod=readonly github.com/gogo/protobuf/protoc-gen-gogo@latest
	go install -mod=readonly github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	go install -mod=readonly github.com/rakyll/statik@latest

install-grpc-env:
	mkdir -p tmp \
	&& wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.5/protoc-3.19.5-linux-x86_64.zip -O tmp/protoc.zip \
	&& unzip -u -d tmp tmp/protoc.zip \
	&& rsync -a tmp/bin/* /usr/local/bin \
	&& rsync -a tmp/include/* /usr/local/include/ \
	&& rm -rf tmp

deepcopy-gen:
	@echo ">> generating apis/${VERSION}/deepcopy_generated.go"
	deepcopy-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	diff ${GOPATH_SRC}/${PROJECT_APIS}/deepcopy_generated.go apis/${VERSION}/deepcopy_generated.go \
	|| mv ${GOPATH_SRC}/${PROJECT_APIS}/deepcopy_generated.go apis/${VERSION}

register-gen:
	@echo ">> generating apis/${VERSION}/zz_generated.register.go"
	register-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	diff ${GOPATH_SRC}/${PROJECT_APIS}/zz_generated.register.go apis/${VERSION}/zz_generated.register.go \
	|| mv ${GOPATH_SRC}/${PROJECT_APIS}/zz_generated.register.go apis/${VERSION}

defaulter-gen:
	@echo ">> generating apis/${VERSION}/zz_generated.defaults.go"
	defaulter-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	diff ${GOPATH_SRC}/${PROJECT_APIS}/zz_generated.defaults.go apis/${VERSION}/zz_generated.defaults.go \
	|| mv ${GOPATH_SRC}/${PROJECT_APIS}/zz_generated.defaults.go apis/${VERSION}

openapi-gen:
	@echo ">> generating apis/${VERSION}/openapi_generated.go"
	openapi-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	diff ${GOPATH_SRC}/${PROJECT_APIS}/openapi_generated.go apis/${VERSION}/openapi_generated.go \
	|| mv ${GOPATH_SRC}/${PROJECT_APIS}/openapi_generated.go apis/${VERSION}

client-gen:
	@echo ">> generating pkg/clients/clientset..."
	rm -rf pkg/clients/clientset
	client-gen --input-dirs ${PROJECT_APIS} \
		--clientset-name='mebius' \
		--fake-clientset=false \
		--input-base=${PROJECT} \
		--input='apis/${VERSION}' \
		--output-package ${CLIENTSET} -h hack.txt \
	--alsologtostderr
	mv ${GOPATH_SRC}/${CLIENTSET} pkg/clients

lister-gen:
	@echo ">> generating pkg/clients/lister..."
	rm -rf pkg/clients/lister
	lister-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${LISTER} -h hack.txt \
	--alsologtostderr
	mv ${GOPATH_SRC}/${LISTER} pkg/clients

informer-gen:
	@echo ">> generating pkg/clients/informer..."
	rm -rf pkg/clients/informer
	informer-gen --input-dirs ${PROJECT_APIS} --versioned-clientset-package ${CLIENTSET}/mebius \
		--output-package ${INFORMER} -h hack.txt \
		--listers-package ${LISTER} \
	--alsologtostderr
	mv ${GOPATH_SRC}/${INFORMER} pkg/clients

go-to-protobuf: vendor
	@echo ">> generating apis/${VERSION}/generated.proto"
	rm -f apis/${VERSION}/generated.proto
	go-to-protobuf --output-base="${GOPATH_SRC}" \
	--apimachinery-packages="-k8s.io/apimachinery/pkg/util/intstr,-k8s.io/apimachinery/pkg/api/resource,-k8s.io/apimachinery/pkg/runtime/schema,-k8s.io/apimachinery/pkg/runtime,-k8s.io/apimachinery/pkg/apis/meta/v1,-k8s.io/api/core/v1" \
	--packages="${PROJECT_APIS},${PROTO_TYPES}" \
	--proto-import "vendor,${GOPATH_SRC}/github.com/gogo/protobuf/protobuf,vendor/k8s.io/apimachinery/pkg/apis/meta/v1" \
	-h hack.txt
	test -f  apis/${VERSION}/generated.proto \
	|| cp ${GOPATH_SRC}/${PROJECT_APIS}/generated.proto apis/${VERSION}
	@echo ">> generating pkg/protos/types/generated.proto"


crd:
	controller-gen crd:crdVersions=v1,allowDangerousTypes=true paths="./apis/..." output:crd:artifacts:config=crds

goimports:
	go install golang.org/x/tools/cmd/goimports@latest

grpc: go-to-protobuf
	protoc \
	-I . \
	-I ${GOPATH_SRC} \
	-I ./vendor \
	-I ${GOPATH_SRC}/github.com/gogo/googleapis \
	-I /usr/local/include \
	--gogo_out=plugins=grpc,paths=source_relative:./ \
	--grpc-gateway_out=logtostderr=true,v=10,allow_patch_feature=true,paths=source_relative:./ \
	--swagger_out=logtostderr=true,v=10:./ \
	--govalidators_out=gogoimport=true,paths=source_relative,:./ \
	pkg/protos/*.proto

fmt: goimports
	go fmt ./...
	${GOPATH}/bin/goimports -l $(shell find . -type f -name '*.go' -not -path "./vendor/*")

vet: 
	go vet ./...

doc:
	cp pkg/protos/*.json docs/static
	statik -m -f -src docs/static

build: clean doc fmt vet
	go build -o bin/mebius main.go

run: fmt vet
	go run ./main.go

test:
	go test ./... -coverprofile cover.out

vendor:
	go mod vendor

install:
	kubectl apply -f crds

uninstall:
	kubectl delete -f crds

clean:
	rm -rf vendor

STAGING_REGISTRY ?= mebius
IMAGE_NAME ?= mebius
TAG ?= latest

IMG ?= ${STAGING_REGISTRY}/${IMAGE_NAME}:${TAG}
docker-build:
	docker buildx build -t ${IMG} . --load

docker-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t ${IMG} . --push