module maistra.io/api

go 1.15

require (
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6
	github.com/maistra/xns-informer v0.0.0-20230111133120-6dfd424e4c7d
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	google.golang.org/protobuf v1.27.1
	k8s.io/api v0.24.1
	k8s.io/apiextensions-apiserver v0.24.1
	k8s.io/apimachinery v0.24.1
	k8s.io/client-go v0.24.1
	k8s.io/code-generator v0.24.1
	sigs.k8s.io/controller-runtime v0.12.1
	sigs.k8s.io/controller-tools v0.7.0
	sigs.k8s.io/yaml v1.3.0
)

replace github.com/go-logr/zapr => github.com/go-logr/zapr v0.2.0
