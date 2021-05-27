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
plural_exceptions        = ServiceExports:ServiceExports,ServiceImports:ServiceImports

# protobuf
out_path = /tmp

gogofast_plugin_prefix := --gogofast_out=plugins=grpc,

comma := ,
empty :=
space := $(empty) $(empty)

importmaps := \
	gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto \
	google/protobuf/any.proto=github.com/gogo/protobuf/types \
	google/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor \
	google/protobuf/duration.proto=github.com/gogo/protobuf/types \
	google/protobuf/struct.proto=github.com/gogo/protobuf/types \
	google/protobuf/timestamp.proto=github.com/gogo/protobuf/types \
	google/protobuf/wrappers.proto=github.com/gogo/protobuf/types \
	google/rpc/status.proto=istio.io/gogo-genproto/googleapis/google/rpc \
	google/rpc/code.proto=istio.io/gogo-genproto/googleapis/google/rpc \
	google/rpc/error_details.proto=istio.io/gogo-genproto/googleapis/google/rpc \
	google/api/field_behavior.proto=istio.io/gogo-genproto/googleapis/google/api \

mapping_with_spaces := $(foreach map,$(importmaps),M$(map),)
gogo_mapping := $(subst $(space),$(empty),$(mapping_with_spaces))

gogofast_plugin := $(gogofast_plugin_prefix)$(gogo_mapping):$(out_path)

protoc = protoc -I. ${gogofast_plugin} --deepcopy_out=${out_path}

security_v1_path := security/v1
security_v1_protos := $(wildcard $(security_v1_path)/*.proto)
security_v1_pb_gos := $(security_v1_protos:.proto=.pb.go)

generate-crd:
	$(CONTROLLER_GEN) crd:preserveUnknownFields=false,crdVersions=v1beta1 paths=$(path_apis) output:dir=./manifests/
	sed -i -e '/---/d' ./manifests/maistra.io_*.yaml

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

remove-proto:
	rm -f ${security_v1_path}/*.gen.go
	rm -f ${security_v1_path}/*.pb.go

generate-proto: remove-proto ${security_v1_pb_gos}

${security_v1_pb_gos}: ${security_v1_protos}
	$(protoc) $^
	cp -r /tmp/maistra.io/api/security/* security

clean:
	rm -rf client manifests
	find core -name zz_generated.deepcopy.go -delete

all: gen
gen: generate-client generate-copy generate-crd generate-proto
	go mod tidy
	go mod vendor

gen-check: clean gen check-clean-repo

check-clean-repo:
	@if [[ -n $$(git status --porcelain) ]]; then echo -e "ERROR: Some files need to be updated, run 'make gen' and include any changed files in your PR. Output:\n"; git status; git diff; exit 1; fi

test:
	go test -race $(path_apis)
