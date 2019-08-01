# Secrets \(soon\)

No matter what kind of application we are running we will most certainly need to manage secrets. 

Kubernetes has a build in way to manage secrets.

I will be honest, its not good nor is it secure. We will see in a second why.

Lets create a secret-

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

