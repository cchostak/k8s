# 1. Create the ssk key-pair if it doesn't exist

```sh
aws ec2 create-key-pair --key-name ec2-ssh-key
```

# 1.1 Deploy through cloudFormation the eks-vpc.yaml

WebConsole

```sh
$ aws cloudformation deploy --template-file ./eks-vpc.yaml --stack-name eks-vpc
```

# 1.2 Create the IAM EKS Role

WebConsole

# 1.3 Create the EKS Cluster in 

WebConsole

# 2. After creating the EKS Cluster through Web Console, configure kubectl locally

aws eks update-kubeconfig --name EKSDeepDive
kubectl config get-contexts
kubectl config set-context EKSDeepDive
kubectl cluster-info

# 2.1 Deploy node Group through CloudFormation

WebConsole

# 3. Opmitimized AMIS

ami-0dc7713312a7ec987 (us-east-1)

# 4. Apply aws-auth-cm 

kubectl apply -f aws-auth-cm.yaml

# 5. Verify it

kubectl get nodes --watch

# 6. Kubernetes Dashboard

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

# 7. Heapter and InfluxDB

kubectl apply -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/heapster.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/influxdb.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/rbac/heapster-rbac.yaml

# 8. Administrative Role and Cluster binding

kubectl apply -f eks-admin-service-account.yaml
kubectl apply -f eks-admin-cluster-role-binding.yaml

# 9. Start Proxy

kubectl proxy --address 0.0.0.0 --accept-hosts '.*' &


# 10. Get a token through IAM authenticator

aws-iam-authenticator -i EKSDeepDive token | jq -r .status.token

# 11. Acess dashboard

http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login

(Paste token)

# 12. Edit/Add new users to admin the cluster

As the new user that wants to join the cluster:

aws sts get-caller-identity

As admin, run:

kubectl edit -n kube-system configmap/aws-auth

# 13. Login to ECR

docker login -u AWS -p $(aws ecr get-login-password --region us-east-1)  https://"$(aws sts get-caller-identity | jq -r .Account)".dkr.ecr.us-east-1.amazonaws.com/

export ECR_REPO_PREFIX=$(aws sts get-caller-identity | jq -r .Account).dkr.ecr.us-east-1.amazonaws.com

# 13.1 Create repos

aws ecr create-repository --repository-name photo-filter
aws ecr create-repository --repository-name photo-storage
aws ecr create-repository --repository-name web-client