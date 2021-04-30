CONTROLLER_GEN = go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go
LISTER_GEN     = go run vendor/k8s.io/code-generator/cmd/lister-gen/main.go
INFORMER_GEN   = go run vendor/k8s.io/code-generator/cmd/informer-gen/main.go
CLIENT_GEN     = go run vendor/k8s.io/code-generator/cmd/client-gen/main.go

empty :=
space := $(empty) $(empty)

kube_api_packages = $(subst $(space),$(empty), \
	$(kube_base_output_package)/core/v1, \
	$(kube_base_output_package)/core/v1alpha1, \
	$(kube_base_output_package)/core/v2 \
	)
kube_base_output_package = maistra.io/api
kube_clientset_package   = $(kube_base_output_package)/client
kube_listers_package     = $(kube_base_output_package)/client/listers
kube_informers_package   = $(kube_base_output_package)/client/informers
path_apis                = "./core/..."
header_file              = "header.go.txt"

generate-crd:
	$(CONTROLLER_GEN) crd:preserveUnknownFields=false,crdVersions=v1beta1 paths=$(path_apis) output:dir=./manifests/
	sed -i -e '/---/d' ./manifests/maistra.io_*.yaml

generate-copy:
	$(CONTROLLER_GEN) object:headerFile=$(header_file) paths=$(path_apis)

generate-client:
	$(CLIENT_GEN) --clientset-name versioned --input-base "" --input $(kube_api_packages) --output-package $(kube_clientset_package) -h $(header_file)
	## Hack - Because we are using core, client-gen hardcodes it to /api
	find client -name core_client.go -exec sed -i 's|config.APIPath = "/api"|config.APIPath = "/apis"|' {} \;

generate-lister:
	$(LISTER_GEN) --input-dirs $(kube_api_packages) --output-package $(kube_listers_package) -h $(header_file)

generate-informer:
	$(INFORMER_GEN) --input-dirs $(kube_api_packages) --versioned-clientset-package $(kube_clientset_package)/versioned --listers-package $(kube_listers_package) --output-package $(kube_informers_package) -h $(header_file)

clean:
	rm -rf client manifests
	find core -name zz_generated.deepcopy.go -delete

all: gen
gen: generate-client generate-lister generate-informer generate-copy generate-crd
	go mod tidy
	go mod vendor

gen-check: clean gen check-clean-repo

check-clean-repo:
	@if [[ -n $$(git status --porcelain) ]]; then echo -e "ERROR: Some files need to be updated, run 'make gen' and include any changed files in your PR. Output:\n"; git status; git diff; exit 1; fi

test:
	go test -race $(path_apis)
