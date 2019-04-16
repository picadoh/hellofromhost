This is a simple web application written in Go that simply prints `Hello from <current host>` whenever the / endpoint is accessed. The goal of this application is for demonstration purposes and it's not intended to be a production-ready service (e.g. Docker, Kubernetes, etc).

## Usage

### Go SDK

    go build hello.go
    ./hello
    curl http://localhost:8080

### Docker

Building the image:

    docker build . -t hellofromhost

Using a pre-built image:

    docker run --name hellofromhost -d -p 8080:8080 picadoh/hellofromhost
    curl http://localhost:8080

### Kubernetes

Deploying:

    kubectl create -f kube.yaml
    kubectl get service hellofromhost
    curl http://<external-ip>

Scaling:

    kubectl scale deployments/hellofromhost --replicas=2

Upgrading:

    kubectl set image deployments/hellofromhost hellofromhost=picadoh/hellofromhost:v2
