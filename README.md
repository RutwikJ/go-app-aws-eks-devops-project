# Go Web App with Docker, Kubernetes, and CI/CD

This project is a small web application built in Go and packaged to demonstrate a practical DevOps workflow. It covers application development, testing, containerization, deployment automation, and cloud-native delivery in a way that is easy to explain in interviews.

## What it demonstrates

- A Go-based web server with static pages
- Automated testing for the main routes
- Containerization with Docker
- Kubernetes deployment manifests
- Helm-based deployment setup
- GitHub Actions CI/CD automation
- Amazon EKS cluster configuration with eksctl

## Architecture

```flowchart LR
A[Developer] --> B[GitHub Repository]
B --> C[GitHub Actions]
C --> D[Docker Image]
D --> E[Kubernetes / Helm]
E --> F[Amazon EKS]
```

## Project structure

- `main.go` - application entry point and route handlers
- `main_test.go` - basic tests for the web routes
- `Dockerfile` - container build steps
- `k8s/manifests/` - Kubernetes deployment files
- `helm/go-web-app-chart/` - Helm chart for deployment
- `.github/workflows/cicd.yaml` - CI/CD pipeline
- `cluster.yaml` - EKS cluster configuration

## Run locally

Make sure Go 1.21 or later is installed.

If you are already inside this project folder, run:

```bash
go run .
```

If you are starting from the repository root, first go into the app folder:

```bash
cd go-web-app
go run .
```

Then open:

```text
http://localhost:8080
```

You can test these routes:

- `/`
- `/about`
- `/contact`

## Build and test

```bash
go test ./...
go build -o go-web-app
```

## Docker

Build the image:

```bash
docker build -t go-web-app .
```

Run the container:

```bash
docker run -p 8080:8080 go-web-app
```

## Kubernetes and Helm

Apply the Kubernetes manifests:

```bash
kubectl apply -f k8s/manifests
```

Or deploy with Helm:

```bash
helm install go-web-app helm/go-web-app-chart
```

## CI/CD pipeline

The GitHub Actions workflow performs the following steps:

- checks out the repository
- sets up Go
- runs build and tests
- runs golangci-lint
- builds and pushes a Docker image to Docker Hub
- updates the Helm chart image tag


