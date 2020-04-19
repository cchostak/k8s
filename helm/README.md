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

```

## Install helm on MAC
```sh
brew install helm
helm repo update
```

## Install on Linux
```sh
kubectl get nodes
curl https://storage.googleapis.com/kubernetes-helm/helm-v2.12.3-linux-amd64.tar.gz > ./helm.tar.gz
tar -xvf ./helm.tar.gz
cd linux-amd64
sudo mv ./helm /usr/local/bin
sudo mv ./tiller  /usr/local/bin
cd
helm version
helm init
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p'{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
kubectl get pods -n kube-system
helm install ./nginx/
curl 10.106.145.215:8888
helm ls --short
helm delete $(helm ls --short)
helm ls
kubectl get pods
```

### helm 3.x.x doesn-t usa tiller anymore

# Creating and pulling charts

```sh
helm create <CHART NAME> 
helm install <CHART NAME> #Will install the boilerplate nginx
helm fetch --untar stable/jenkins #Downloads the chart in .tgz
```

# Charts explained

```sh
wordpress/
  Chart.yaml          # A YAML file containing information about the chart
  LICENSE             # OPTIONAL: A plain text file containing the license for the chart
  README.md           # OPTIONAL: A human-readable README file
  values.yaml         # The default configuration values for this chart
  values.schema.json  # OPTIONAL: A JSON Schema for imposing a structure on the values.yaml file
  charts/             # A directory containing any charts upon which this chart depends.
  crds/               # Custom Resource Definitions
  templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
  templates/NOTES.txt # OPTIONAL: A plain text file containing short usage notes
```

Helm reserves use of the charts/, crds/, and templates/ directories, and of the listed file names. Other files will be left as they are.

# Helm completion

helm completion zsh >> ~/.zshrc

# Generating a custom values file 

This is useful when applying the same Chart to multiple environments.

```sh
grep -v '#' values.yaml > custom.yaml
helm install --values=./custom.yaml stable/wordpress --generate-name
```

# Set up a local chart repository directory

```sh
helm fetch --untar stable/jenkins #Downloads the chart in .tgz and extract it
```

Change the contents of the Chart as you wish.

```sh
helm package <FOLDER>
helm repo index .
helm repo add <A NAME> <URL POINTING TO THE PUBLIC FOLDER WHERE THE INDEX IS LOCATED>
helm search <YOUR CHART>
```

# Releases

```sh
helm fetch --untar stable/jenkins #Downloads the chart in .tgz and extract it
```

Change the contents of the Chart as you wish. Upgrade Version

```sh
helm package <FOLDER>
helm repo index .
helm repo add <A NAME> <URL POINTING TO THE PUBLIC FOLDER WHERE THE INDEX IS LOCATED>
helm repo update
helm search <YOUR CHART>
helm install --version=5.2.2 <repp>/<package>

kubectl get sts # Stateful Set, this is destructive
kubectl delete sts <name of the stateful set>
helm ls --short
helm upgrade --recreate-pods <name of the release from above> <repo>/<package> 
helm history <release name>
helm rollback <release name> <revision number>
```

