PROJECT=github.com/LTitan/Mebius/pkg/apis

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
	go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1
	go install github.com/gogo/protobuf/protoc-gen-gogo@v1.3.2

deepcopy-gen:
	deepcopy-gen --input-dirs ${PROJECT} \
		--output-package ${PROJECT} -h hack.txt \
	--alsologtostderr

register-gen:
	register-gen --input-dirs ${PROJECT} \
		--output-package ${PROJECT} -h hack.txt \
	--alsologtostderr

defaulter-gen:
	defaulter-gen --input-dirs ${PROJECT} \
		--output-package ${PROJECT} -h hack.txt \
	--alsologtostderr

openapi-gen:
	openapi-gen --input-dirs ${PROJECT} \
		--output-package ${PROJECT} -h hack.txt \
	--alsologtostderr

client-gen:
	rm -rf pkg/clients/clientset
	client-gen --input-dirs ${PROJECT} \
		--clientset-name='mebius' \
		--fake-clientset=false \
		--input-base='github.com/LTitan/Mebius' \
		--input='pkg/apis' \
		--output-package github.com/LTitan/Mebius/pkg/clients/clientset -h hack.txt \
	--alsologtostderr

lister-gen:
	rm -rf pkg/clients/lister
	lister-gen --input-dirs ${PROJECT} \
		--output-package github.com/LTitan/Mebius/pkg/clients/lister -h hack.txt \
	--alsologtostderr

informer-gen:
	rm -rf pkg/clients/informer
	informer-gen --input-dirs ${PROJECT} \
		--output-package github.com/LTitan/Mebius/pkg/clients/informer -h hack.txt \
	--alsologtostderr

go-to-protobuf:
	go-to-protobuf --output-base="${GOPATH}/src" --packages="${PROJECT}" -h hack.txt

crd:
	controller-gen crd:crdVersions=v1,trivialVersions=true,allowDangerousTypes=true paths="./pkg/apis/..." output:crd:artifacts:config=crds