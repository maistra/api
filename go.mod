module maistra.io/api

go 1.15

require (
	github.com/ghodss/yaml v1.0.0
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/code-generator v0.20.2
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/controller-tools v0.5.0
	sigs.k8s.io/yaml v1.2.0
)

replace k8s.io/client-go => k8s.io/client-go v0.19.3

replace k8s.io/code-generator => k8s.io/code-generator v0.19.3

replace k8s.io/api => k8s.io/api v0.19.3

replace k8s.io/apimachinery => k8s.io/apimachinery v0.19.3
