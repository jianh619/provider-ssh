# provider-ssh

`provider-ssh` is a minimal [Crossplane](https://crossplane.io/) Provider
, which will help connecting to remote vm and do something .


## Getting Started

### Prerequisites

- You need install a kubernetes cluster https://kubernetes.io/
- You need install crossplane https://crossplane.io

### Play

- Clone this repository :

```console
git clone https://github.com/jianh619/provider-ssh.git
```

- Run against a Kubernetes cluster:

```console
make run
```

- Apply providerConfig 
  
  Edit examples/provider/config.yaml
  
```console
 apiVersion: ssh.crossplane.io/v1alpha1
 kind: ProviderConfig
 metadata:
   name: example
 spec:
   ip: <ip>
   user: <user>
   password: <password>
   credentials:
     source: Secret
     secretRef:
       namespace: crossplane-system
       name: example-provider-secret
       key: credentials
```
**Note:** You need speicify ip/user/password for connecting this vm :

```
kubectl apply -f examples/provider/config.yaml
```

- Create a File Kind Customer Resource

```console
kubectl apply -f examples/sample/file.yaml
```

It will make sure there's a file named `ssh-test` in your remote VM (/root/ssh-test)

## Developing

Run against a Kubernetes cluster:

```console
make run
```

Run against a Kubernetes cluster:
```console
make run
```

Build, push, and install:

```console
make all
```

Build image:

```console
make image
```

Push image:

```console
make push
```

Build binary:

```console
make build
```
