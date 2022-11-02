CONTROLLER_GEN   = go run sigs.k8s.io/controller-tools/cmd/controller-gen
LISTER_GEN       = go run k8s.io/code-generator/cmd/lister-gen
INFORMER_GEN     = go run k8s.io/code-generator/cmd/informer-gen
CLIENT_GEN       = go run k8s.io/code-generator/cmd/client-gen
CONVERSION_GEN   = go run k8s.io/code-generator/cmd/conversion-gen
DOC_GEN          = go run ./tools/cmd/doc
XNS_INFORMER_GEN = go run github.com/maistra/xns-informer/cmd/xns-informer-gen

empty :=
space := $(empty) $(empty)

kube_api_packages = $(subst $(space),$(empty), \
	$(kube_base_output_package)/core/v1, \
	$(kube_base_output_package)/core/v1alpha1, \
	$(kube_base_output_package)/core/v2, \
	$(kube_base_output_package)/federation/v1 \
	)
kube_base_output_package = maistra.io/api
kube_clientset_package   = $(kube_base_output_package)/client
kube_listers_package     = $(kube_base_output_package)/client/listers
kube_informers_package   = $(kube_base_output_package)/client/informers
xns_informers_package    = $(kube_base_output_package)/client/xnsinformer
path_apis                = "./core/...;./federation/..."
header_file              = "header.go.txt"
plural_exceptions        = ExportedServices:ExportedServices,ImportedServices:ImportedServices

security_v1_path := security/v1
security_v1_protos := $(wildcard $(security_v1_path)/*.proto)
security_v1_pb_gos := $(security_v1_protos:.proto=.pb.go)

docs_path := docs/crd

generate-crd:
	$(CONTROLLER_GEN) crd:crdVersions=v1 paths=$(path_apis) output:dir=./manifests/
	sed -i -e '/---/d' ./manifests/maistra.io_*.yaml ./manifests/federation.maistra.io_*.yaml

generate-docs:
	rm -rf $(docs_path)
	$(DOC_GEN) paths="maistra.io/api/core/...;maistra.io/api/federation/..." output:dir=$(docs_path) doc:format=adoc,depth=2

generate-copy:
	$(CONTROLLER_GEN) object:headerFile=$(header_file) paths=$(path_apis)

generate-client:
	rm -rf maistra.io
	$(CLIENT_GEN) --clientset-name versioned --input-base "" --input $(kube_api_packages) --output-package $(kube_clientset_package) -h $(header_file) --output-base ./ --plural-exceptions $(plural_exceptions)
	$(LISTER_GEN) --input-dirs $(kube_api_packages) --output-package $(kube_listers_package) -h $(header_file) --output-base ./ --plural-exceptions $(plural_exceptions)
	$(INFORMER_GEN) --input-dirs $(kube_api_packages) --versioned-clientset-package $(kube_clientset_package)/versioned --listers-package $(kube_listers_package) --output-package $(kube_informers_package) -h $(header_file) --output-base ./ --plural-exceptions $(plural_exceptions)
	rm -rf client && mv maistra.io/api/client . && rm -rf maistra.io
	## Hack - Because we are using core, client-gen hardcodes it to /api
	find client -name core_client.go -exec sed -i 's|config.APIPath = "/api"|config.APIPath = "/apis"|' {} \;
	## Hack - Because we are not using +groupName=maistra.io in doc.go, the incorrect group is used in fake types
	find client/versioned/typed/core -name "*.go" -exec sed -i 's|Group: ""|Group: "maistra.io"|' {} \;
	$(XNS_INFORMER_GEN) --output-base ./ --output-package $(xns_informers_package) --single-directory --input-dirs $(kube_api_packages) --versioned-clientset-package $(kube_clientset_package)/versioned --listers-package $(kube_listers_package) -h $(header_file) --plural-exceptions $(plural_exceptions)
	rm -rf client/xnsinformer && mv maistra.io/api/client/xnsinformer ./client && rm -rf maistra.io

remove-proto:
	rm -f ${security_v1_path}/*.gen.go
	rm -f ${security_v1_path}/*.pb.go

generate-proto: remove-proto ${security_v1_pb_gos}

${security_v1_pb_gos}: ${security_v1_protos}
	buf generate --path ${security_v1_path}

generate-conversion:
	# --base-peer-dirs includes metav1, which causes conversion funcs to be generated for converting metav1.Condition to/from v1.Condition
	$(CONVERSION_GEN) --input-dirs $(kube_api_packages) --output-package $(kube_base_output_package) -h $(header_file) --output-base ./ --base-peer-dirs k8s.io/apimachinery/pkg/conversion,k8s.io/apimachinery/pkg/runtime
	mv maistra.io/api/core/v1alpha1/* core/v1alpha1/ && rm -rf maistra.io

clean:
	rm -rf client manifests/*.yaml
	find core -name zz_generated.deepcopy.go -delete -o -name conversion_generated.go -delete
	find federation -name zz_generated.deepcopy.go -delete -o -name conversion_generated.go -delete

all: gen
gen: clean generate-copy generate-client generate-crd generate-conversion generate-proto generate-docs
	go mod tidy

gen-check: gen check-clean-repo

check-clean-repo:
	@if [[ -n $$(git status --porcelain) ]]; then echo -e "ERROR: Some files need to be updated, run 'make gen' and include any changed files in your PR. Output:\n"; git status; git diff; exit 1; fi

test:
	go test -race ./core/... ./federation/...
