* Instsall and configure Minikube DONE
* Instsall and configure Kubectl DONE
* Instsall and configure helm DONE

* Install and configure CI/CD ------>> CHANGING to BUILDBOT (Easier to install and use for the purpose of this project.)
# Reference: https://github.com/kubernetes/charts/tree/master/stable/concourse
# Command: 
# helm install --name concourse --set web.service.type=NodePort  stable/concourse
# * Configure money-transfer-simulation pipeline DONE

* Store image in DockerHub DONE

* Use Helm to install service. DONE.

# Helm command to meke it work in kubernetes.
# helm install --name=users -- ../../charts/users

# while :; do; sleep 1; curl -i http://192.168.99.100:30080/api/v1/users/1; done;

helm install --wait --timeout 30 --set image.tag=$tag --set serviceInfo.name=$_serviceName --set service.containerPort=80 --set service.port=$_port --name=web -- $pathDir/../../charts/deployment-service

helm upgrade ${_serviceName} --wait --timeout 30 --set image.repository=$_serviceImageName --set image.tag=$tag --set serviceInfo.name=$_serviceName  -- $pathDir/../../charts/deployment-service

* Split services in user create/autenticate, transference.
* Testing Services
* Update new image automatically

* Add bots to make random operations with service
# For creting users. 
# curl -i https://randomuser.me/api/\?inc=name,mail,phone
* Leave bug in code to kill pod. DONE.
* * Fix it.  DONE.
* Leave bug in code to use 100% CPU. DONE.
* * Fix it. DONE.


# ------------------------------------------------------------------------------------------

* Linkerd ??
# https://linkerd.io/getting-started/k8s/
# https://github.com/linkerd/linkerd-examples/issues/130
# https://github.com/linkerd/linkerd-examples/tree/master/k8s-daemonset

* Istio ??
# https://meteatamel.wordpress.com/2018/04/24/istio-101-with-minikube/

* Dashboards
* * Grafana
* * Prometheus
* * Linkerd / Istio
* * EFK NOT POSIBLE -> Too much resources required for it to work properlly.

Articles:
# https://dzone.com/articles/microservices-the-good-the-bad-and-the-ugly
# https://medium.com/@zahirkoradia/microservices-the-good-the-bad-the-ugly-a2f09248092d

