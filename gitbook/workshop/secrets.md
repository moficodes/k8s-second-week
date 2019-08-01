# Secrets

No matter what kind of application we are running we will most certainly need to manage secrets. 

Kubernetes has a build in way to manage secrets.

I will be honest, its not good nor is it secure. We will see in a second why.

## Creating Secret

### Manually

Lets say we have an application that has super secret password and it is `super-secret-password` 

```text
echo -n 'super-secret-password' | base64
```

You should get 

`c3VwZXItc2VjcmV0LXBhc3N3b3Jk`

 Lets create the yaml file for this secret. The file is located at `k8s/secrets/secret-manual.yaml`

```text
kubectl apply -f k8s/secrets/secret-manual.yaml
```

If you want to see all the secrets run

```text
kubectl get secrets
```

You will see bunch of secrets printed

The cluster adds some secrets by default for using private registry under the same account.

```text
default-au-icr-io     kubernetes.io/dockerconfigjson        1      7h46m
default-de-icr-io     kubernetes.io/dockerconfigjson        1      7h47m
default-icr-io        kubernetes.io/dockerconfigjson        1      7h47m
default-jp-icr-io     kubernetes.io/dockerconfigjson        1      7h46m
default-token-5tdlr   kubernetes.io/service-account-token   3      7h49m
default-uk-icr-io     kubernetes.io/dockerconfigjson        1      7h47m
default-us-icr-io     kubernetes.io/dockerconfigjson        1      7h47m
kube201-workshop01    Opaque                                2      7h29m
password              Opaque                                1      3m16s
```

But our password secret is also there. Great success!

How can we see it.

Lets describe said resource

```text
kubectl describe secret password
```

The kubectl command helpfully hide

```text
Name:         password
Namespace:    default
Labels:       <none>
Annotations:
Type:         Opaque

Data
====
password:  21 bytes
```

But this is not secure at all. We can just see it by looking at the yaml representation of the secret.

```text
kubectl get secret password -o yaml
```

This prints

```text
apiVersion: v1
data:
  password: c3VwZXItc2VjcmV0LXBhc3N3b3Jk <------ Oh noooo!!!
kind: Secret
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","data":{"password":"c3VwZXItc2VjcmV0LXBhc3N3b3Jk"},"kind":"Secret","metadata":{"annotations":{},"name":"password","namespace":"default"},"type":"Opaque"}
  creationTimestamp: "2019-08-01T05:08:25Z"
  name: password
  namespace: default
  resourceVersion: "9247607"
  selfLink: /api/v1/namespaces/default/secrets/password
  uid: 646fd450-b41a-11e9-8ea5-96001dc10e67
type: Opaque
```

We can just copy that string and decode that and get the password in text. 

Or we can do that with a little bit of linux magic

```text
kubectl get secrets password -o jsonpath='{.data.password}' | base64 --decode
```

> Using jsonpath we are extracting the password from the json output that piping that to base64 decode.

And with that we get our `super-secret-password` back. Not so secure.

### Using Kubectl

We can also create secret directly using `kubectl` 

Run

```text
kubectl create secret generic password-kubectl \
    --from-literal=password=password-from-literal
```

Different method same out come.

## Consuming Secret

Creating a secret means nothing if we can't use it. 

There are basically two ways you can consume secret in a kubernetes environment.

1. Passed as an environmental variable
2. Mounted a volume and consumed as a file

### Environmental Variable

```text
env:
- name: SECRET_USERNAME
   valueFrom:
      secretKeyRef:
         name: mysecret
         key: passeword
```

The issue with this approach is that kubernetes does not do any kind of  differentiation between a password env variable and a regular environmental variable. So if your application does log out things anywhere it often writes the env variable in case of a crash. So you will leak your own secrets.

### Volume Mount

```text
spec:
   volumes:
      - name: "secretstest"
         secret:
            secretName: password
   containers:
      - image: nginx
         name: awebserver
         volumeMounts:
            - mountPath: "/tmp/mysec"
            name: "secretstest"
```

This is a bit more secure than env vars. But if someone gets access to the running pod, they can literally read the file. We need to make our images more secure, but thats a separate workshop on its own.

## What do we do?

So the secrets on kubernetes is bad, I hope we all see that. 

Then what do we do?

![](../.gitbook/assets/download.gif)



We could encrypt our own secrets, use encrypted inter cluster communication for the time we are passing secrets. We can encrypt our etcd cluster which is the underlying storage for kubernetes.

The major cloud providers have key management systems that can be used in this way. Other third-party solutions include [HashiCorp Vault](https://www.vaultproject.io/) and [CyberArk Conjur](https://www.conjur.org/).



