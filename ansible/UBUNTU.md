Get the Docker gpg key:
```sh

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```
Add the Docker repository:
```sh

sudo add-apt-repository    "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
```
Get the Kubernetes gpg key:
```sh

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
```
Add the Kubernetes repository:
```sh

cat << EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
```
Update your packages:
```sh

sudo apt-get update
```
Install Docker, kubelet, kubeadm, and kubectl:
```sh

sudo apt-get install -y docker-ce=18.06.1~ce~3-0~ubuntu kubelet=1.15.7-00 kubeadm=1.15.7-00 kubectl=1.15.7-00
```
Hold them at the current version:
```sh

sudo apt-mark hold docker-ce kubelet kubeadm kubectl
```
Add the iptables rule to sysctl.conf:
```sh

echo "net.bridge.bridge-nf-call-iptables=1" | sudo tee -a /etc/sysctl.conf
```
Enable iptables immediately:
```sh

sudo sysctl -p
```
Initialize the cluster (run only on the master):
```sh

sudo kubeadm init --pod-network-cidr=10.244.0.0/16
```
Set up local kubeconfig:
```sh

mkdir -p $HOME/.kube

sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config

sudo chown $(id -u):$(id -g) $HOME/.kube/config
```
Apply Flannel CNI network overlay:
```sh

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```
Join the worker nodes to the cluster:
```sh

sudo kubeadm join [your unique string from the kubeadm init command]
```
Verify the worker nodes have joined the cluster successfully:
```sh

kubectl get nodes
```
Compare this result of the kubectl get nodes command:

NAME                            STATUS   ROLES    AGE   VERSION
chadcrowell1c.mylabserver.com   Ready    master   4m18s v1.15.7
chadcrowell2c.mylabserver.com   Ready    none     82s   v1.15.7
chadcrowell3c.mylabserver.com   Ready    none     69s   v1.15.7