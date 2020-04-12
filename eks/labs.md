# LAB 1

## Configure your CLI

aws configure

## Create EKS cluster using provided cli tool ekscluster https://eksctl.io/introduction/installation/

eksctl create cluster --region=us-east-1 --node-type=t2.small

or 

eksctl create cluster -f eksctl/cluster.yaml

Once you have created a cluster, you will find that cluster credentials were added in ~/.kube/config. If you have kubectl v1.10.x as well as aws-iam-authenticator commands in your PATH, you should be able to use kubectl

# LAB 2

```sh
aws configure

aws eks list-clusters

aws eks update-kubeconfig --name wonderful-badger-1586708941

kubectl get nodes

git clone https://github.com/linuxacademy/eks-deep-dive-2019.git

cd eks-deep-dive-2019/2-LA1-Deploying-an-Application-to-EKS


kubectlapply -f deployment.yaml
ubectl apply -f service.yaml

kubectl get all -o wide

curl ac0f9ae137cef11ea97c40ab4a9cabcc-1878647516.us-east-1.elb.amazonaws.com -v
```

# LAB 3

Navigate to the AWS Management Console
Under EC2, find Autoscaling Groups
Select the autoscaling group already created
Under Actions > Edit, change min=2 and max=8
Click Save

Edit cluster_autoscaler.yaml, replacing <AUTOSCALING GROUP NAME> with the autoscaling group name you found in the console.

Copy the content of asg-policy.json
Navigate to IAM and find the role corresponding to the EC2 worker node group
Add an inline policy
Paste the content of asg-policy.json
Save


```sh
vim cluster_autoscaler.yaml 
cat asg-policy.json 
kubectl apply -f cluster_autoscaler.yaml 
kubectl apply -f nginx.yaml 
kubectl get all -o wide
kubectl get pods
kubectl get rs
kubectl scale deployment nginx-scaleout --replicas=10
kubectl get all -o wide
kubectl expose deployment nginx-scaleout --port 80 --type LoadBalancer
kubectl get services -w
curl a3eb2a5dd7cf211ea9f440ae3f7962e2-2132252827.us-east-1.elb.amazonaws.com -v
while true; do curl a3eb2a5dd7cf211ea9f440ae3f7962e2-2132252827.us-east-1.elb.amazonaws.com -v; done
```

