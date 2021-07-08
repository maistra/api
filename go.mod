module maistra.io/api

go 1.15

require (
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.1
	github.com/google/go-cmp v0.5.2
	github.com/maistra/xns-informer v0.0.0-20210707160032-977ec17e2e0e
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	google.golang.org/grpc v1.28.1
	k8s.io/api v0.20.1
	k8s.io/apiextensions-apiserver v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
	k8s.io/code-generator v0.20.1
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/controller-tools v0.4.1
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/go-logr/zapr => github.com/go-logr/zapr v0.2.0
