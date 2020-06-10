# K8S Autoscaling go application
This repo has a autoscaling go webserver running with [Kubernetes](https://kubernetes.io/pt/).

A go server containing the deployment and cluster ip service for Kubernetes. The docker image generated in this app is passed by a CI pipeline in the [Cloud Builder](https://cloud.google.com/cloud-build) and if its everything fine with the image then the Cloud Builder pushed the image into [dockerhub repo](https://hub.docker.com/repository/docker/leorrodrigues/go-hpa).

To ensure that the website will be available in view of the uncertainty of users access, the server will increase automatically his number of pods when the current cpu pods percentage hit a specific pre-defined value.

## Usage

To use these apps you need the [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/) installed in your own machine.

To execute the Go webserver:
```bash
kubectl apply -f devops/service.yaml
kubectl apply -f devops/deployment.yaml
kubectl apply -f devops/hpa.yaml
```

To watch the kubernetes autoscaling the server, you can use:
```bash
watch kubectl get hpa
```

To simulate users access into the server, just run a pod inside kubernetes and use wget:
```bash
kubectl run -it $POD_NAME --image=busybox

while true; do wget -qO- http://go-hpa.default.svc.cluster.local; done;
```
