# Helm Deep Dive

Kubernetes Package manager

Helm is for kubernetes what yum is for RedHat.

Provides Automated installation, version control, dependency management and automated removal

```sh
apt install -y docker.io
systemctl enable docker
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubeadm=1.13\* kubectl=1.13\* kubelet=1.13\* kubernetes-cni=0.7\*


kubeadm init --kubernetes-version stable-1.13 --token-ttl 0 --pod-network-cidr=10.244.0.0/16


kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

git clone https://github.com/linuxacademy/content-kubernetes-helm.git ./rook
cd ./rook/cluster/examples/kubernetes/ceph


kubectl create -f operator.yaml

kubectl create -f cluster.yaml

kubectl get pods -n rook-ceph -w

kubectl create -f storageclass.yaml

brew install helm
helm repo update
```

helm 3.x.x doesn-t usa tiller anymore