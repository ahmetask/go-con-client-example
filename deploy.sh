docker build -t go-con-client .
kubectl delete deployment go-con-client
kubectl delete service go-con-client
kubectl apply -f deployment.yml
kubectl apply -f service.yml