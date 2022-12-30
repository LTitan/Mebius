PROJECT=github.com/LTitan/Mebius
VERSION=v1alpha1
PROJECT_APIS=${PROJECT}/pkg/apis/${VERSION}
CLIENTSET=${PROJECT}/pkg/clients/clientset
INFORMER=${PROJECT}/pkg/clients/informer
LISTER=${PROJECT}/pkg/clients/lister

ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
	export GOPATH
endif
GOPATH_SRC=${GOPATH}/src

all: register-gen deepcopy-gen defaulter-gen openapi-gen client-gen lister-gen informer-gen

install-tools:
	go install k8s.io/code-generator/cmd/go-to-protobuf@v0.25.3
	go install k8s.io/code-generator/cmd/client-gen@v0.25.3
	go install k8s.io/code-generator/cmd/informer-gen@v0.25.3
	go install k8s.io/code-generator/cmd/deepcopy-gen@v0.25.3
	go install k8s.io/code-generator/cmd/lister-gen@v0.25.3
	go install k8s.io/code-generator/cmd/register-gen@v0.25.3
	go install k8s.io/code-generator/cmd/openapi-gen@v0.25.3
	go install k8s.io/code-generator/cmd/defaulter-gen@v0.25.3
	go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.10.0
	go install github.com/gogo/protobuf/protoc-gen-gogo@v1.3.2

deepcopy-gen:
	@echo ">> generating pkg/apis/deepcopy_generated.go"
	deepcopy-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	mv ${GOPATH_SRC}/${PROJECT_APIS}/deepcopy_generated.go pkg/apis/${VERSION}

register-gen:
	@echo ">> generating pkg/apis/zz_generated.register.go"
	register-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	mv ${GOPATH_SRC}/${PROJECT_APIS}/zz_generated.register.go pkg/apis/${VERSION}

defaulter-gen:
	@echo ">> generating pkg/apis/zz_generated.defaults.go"
	defaulter-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	mv ${GOPATH_SRC}/${PROJECT_APIS}/zz_generated.defaults.go pkg/apis/${VERSION}

openapi-gen:
	@echo ">> generating pkg/apis/openapi_generated.go"
	openapi-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr
	mv ${GOPATH_SRC}/${PROJECT_APIS}/openapi_generated.go pkg/apis/${VERSION}

client-gen:
	@echo ">> generating pkg/clients/clientset..."
	rm -rf pkg/clients/clientset
	client-gen --input-dirs ${PROJECT_APIS} \
		--clientset-name='mebius' \
		--fake-clientset=false \
		--input-base=${PROJECT} \
		--input='pkg/apis/${VERSION}' \
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

go-to-protobuf:
	go-to-protobuf --output-base="${GOPATH}/src" --packages="${PROJECT_APIS}" -h hack.txt

crd:
	controller-gen crd:crdVersions=v1,allowDangerousTypes=true paths="./pkg/apis/${VERSION}..." output:crd:artifacts:config=crds

goimports:
	go install golang.org/x/tools/cmd/goimports@latest

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...