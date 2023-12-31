# go-expert
Golang repository to study basic and advanced concepts of the language

##

**Package manager**

Module initializing:
```
$ go mod init
```

To install package:
```
$ go get -u package-name-url
```

Packaging update and control of version in go.mod:
```
$ go mod tidy
```

Replace package module in dev environment (workaround):
```
$ go mod edit -replace package-name-url=internal-package-path
$ go mod tidy
```

Replace package module in dev environment using workspace (correct approach):
```
$ go work init internal-package-path-1 internal-package-path-2
$ go mod tidy -e
```

##

**Running unit tests**

Running tests generating code coverage information in verbose mode
```
$ go test -coverprofile=coverage.out -v
```

Visualizing where needs to cover in HTML format
```
$ go tool cover -html=coverage.out
```

Benchmark visualization
```
$ go test -bench=. -run=^# -benchmem
```

Performing test with Fuzzing
```
$ go test -fuzz=. -run=^# -fuzztime=5s
```

##
**Performing build of application**

```
MacOS
$ GOOS=darwin GOARCH=os_arch_name go build -o service-name

Linux
$ GOOS=linux GOARCH=os_arch_name go build -o service-name

Windows
$ GOOS=windows GOARCH=os_arch_name go build -o service-name
```

To see all supported platforms
```
$ go tool dist list
```

Need help ?
```
$ go help
```

##
**Kubernetes Owerview**

`considering you have already configured the deployment.yaml and service.yaml`

Create a kubernetes cluster using Kind tool
```
$ kind create cluster --name=some-name-here
```

To see the cluster
```
$ kubectl get nodes
```

To delete a cluster
```
$ kind delete cluster --name=cluster-name
```

Up the pods (according to your replicas configured on deployment.yaml)
```
$ kubectl apply -f deployment.yaml
```

To see all pods upper
```
$ kubectl get pods
```

To delete pods
```
$ kubectl delete pods service-name
```

Up the services to enable the access between pods (according to your service.yaml)
```
$ kubectl apply -f service.yaml
```

To see all services upper
```
$ kubectl get svc
```

To delete services
```
$ kubectl delete svc service-name
```

Forwarding the access for your localhost
```
$ kubectl port-forward svc/serverscv port:port

"serverscv" and the ports was declared in service.yaml
```

##

**Useful links**
- [Convert Json data to Go Struct](https://transform.tools/json-to-go)
- [Go Documentation](https://go.dev/doc/)
- [Go Http](https://pkg.go.dev/net/http)
- [Go Http NewRequest](https://pkg.go.dev/net/http#NewRequest)
- [Go Http Status Code](https://go.dev/src/net/http/status.go)
- [Go ServeMux](https://pkg.go.dev/net/http#ServeMux)
- [Go Text Templates](https://pkg.go.dev/text/template)
- [Go Html Templates](https://pkg.go.dev/html/template)
- [Go Context](https://pkg.go.dev/context)
- [Go UUID](https://pkg.go.dev/github.com/google/uuid)
- [Go SQL](https://pkg.go.dev/database/sql)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)
- [Go ORM](https://gorm.io/docs/index.html)
- [Go Testing](https://pkg.go.dev/testing)
- [Go Testify](https://github.com/stretchr/testify)
- [Viper](https://github.com/spf13/viper)
- [Go-Chi JWT](https://github.com/go-chi/jwtauth)
- [Go BCrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Kubernetes](https://kubernetes.io/docs/tasks/tools/)
- [Kind K8S](https://kind.sigs.k8s.io/)


##

**Useful api links**
- [Via CEP](https://viacep.com.br/)
