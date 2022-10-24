PROJECT=github.com/LTitan/Mebius
PROJECT_APIS=${PROJECT}/pkg/apis
CLIENTSET=${PROJECT}/pkg/clients/clientset
INFORMER=${PROJECT}/pkg/clients/informer
LISTER=${PROJECT}/pkg/clients/lister

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
	deepcopy-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr

register-gen:
	register-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr

defaulter-gen:
	defaulter-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr

openapi-gen:
	openapi-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${PROJECT_APIS} -h hack.txt \
	--alsologtostderr

client-gen:
	rm -rf pkg/clients/clientset
	client-gen --input-dirs ${PROJECT_APIS} \
		--clientset-name='mebius' \
		--fake-clientset=false \
		--input-base=${PROJECT} \
		--input='pkg/apis' \
		--output-package ${CLIENTSET} -h hack.txt \
	--alsologtostderr

lister-gen:
	rm -rf pkg/clients/lister
	lister-gen --input-dirs ${PROJECT_APIS} \
		--output-package ${LISTER} -h hack.txt \
	--alsologtostderr

informer-gen:
	rm -rf pkg/clients/informer
	informer-gen --input-dirs ${PROJECT_APIS} --versioned-clientset-package ${CLIENTSET}/mebius \
		--output-package ${INFORMER} -h hack.txt \
		--listers-package ${LISTER} \
	--alsologtostderr

go-to-protobuf:
	go-to-protobuf --output-base="${GOPATH}/src" --packages="${PROJECT_APIS}" -h hack.txt

crd:
	controller-gen crd:crdVersions=v1,allowDangerousTypes=true paths="./pkg/apis/..." output:crd:artifacts:config=crds

goimports:
	go install golang.org/x/tools/cmd/goimports@latest