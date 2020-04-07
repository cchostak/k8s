# After creating the EKS Cluster through Web Console, configure kubectl locally

aws eks update-kubeconfig --name EKSDeepDive
kubectl config get-contexts
kubectl config set-context EKSDeepDive
kubectl cluster-info

# Opmitimized AMIS

ami-0dc7713312a7ec987 (us-east-1)

# Apply aws-auth-cm 

kubectl apply -f aws-auth-cm.yaml

# Verify it

kubectl get nodes --watch

# Kubernetes Dashboard

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

# Heapter and InfluxDB

kubectl apply -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/heapster.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/influxdb.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/rbac/heapster-rbac.yaml

# Administrative Role and Cluster binding

kubectl apply -f eks-admin-service-account.yaml
kubectl apply -f eks-admin-cluster-role-binding.yaml

# Start Proxy

kubectl proxy --address 0.0.0.0 --accept-hosts '.*' &


# Get a token through IAM authenticator

aws-iam-authenticator -i EKSDeepDive token | jq -r .status.token

# Acess dashboard

http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login

(Paste token)

# Edit/Add new users to admin the cluster

As the new user that wants to join the cluster:

aws sts get-caller-identity

As admin, run:

kubectl edit -n kube-system configmap/aws-auth