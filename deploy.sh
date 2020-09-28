docker build -t go-con-client .
kubectl delete deployment go-con-client
kubectl delete service go-con-client
kubectl create -f deployment.yml
kubectl create -f service.yml
