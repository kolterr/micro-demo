module micro-demo

go 1.13

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0

	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190923035154-9ee001bba392

	golang.org/x/net => github.com/golang/net v0.0.0-20190923162816-aa69164e4478

	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e

	golang.org/x/sys => github.com/golang/sys v0.0.0-20190924154521-2837fb4f24fe

	golang.org/x/text => github.com/golang/text v0.3.2
)

require (
	github.com/Azure/go-autorest/autorest v0.9.1 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.6.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/hashicorp/consul v1.6.1
	github.com/hashicorp/consul/api v1.2.0
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.3.0
)
