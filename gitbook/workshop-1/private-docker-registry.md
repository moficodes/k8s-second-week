# Private Docker Registry

Docker Hub Public Images work pretty great. But when we are talking about enterprise software we want to keep the code secure. We also need to keep our images secure, because anyone with access to the docker image essentially has the business logic of your application.

Docker hub is where almost all public images are stored. Which is where we have been getting our images from so far. But for production code its neither secure nor acceptable to put our production code in a public docker registry.

Almost all major cloud provider and some image registry provider including docker gives us a private docker registry. The usage of private docker registry works pretty much the same way. 

We will use ICR or IBM Container Registry that comes built in with your IBM Cloud account.

## Building a Image and Pushing to ICR

So far we have been using things under the IBM account. That account already has the docker secrets set up. But lets use our own account to use as an image registry.

We can \(re\) login from the cloudshell

```text
ibmcloud login --sso
```

This would ask to get one time code from a url go to that url copy and paste that code. It will list 2 account. Choose the one with your name. \(Number 1\)

> _**You wont see any text when you paste, so don't paste multiple times.**_

Login to the container registry

```text
ibmcloud cr login
```

Lets than create a namespace

```text
ibmcloud cr namespace-add <unique-name>
```

Lets build the `os-signal` image and push it to our private registry.

Also to help with copy-paste export your namespace name as a path variable

```text
export NAMESPACE="<unique-name>"
```

```text
ibmcloud cr build -t us.icr.io/$NAMESPACE/os-signal:0.0.1 src/os-signal/
```

`ibmcloud cr build` is a wrapper around docker build that does docker build and push in a single command and it also doesn't store the image in the local system. Pretty great for a cloud environment.

Lets see if our image got the right place.

```text
ibmcloud cr images
```

It should return the image that we just created and pushed.

## Use Image in Kubernetes

Lets try using this image in our deployment.

But first login to the ibm account.

```text
ibmcloud login --sso
```

This time choose account 2. \(the one with IBM\)

Update the file at `k8s/private-docker-registry/deployment-private.yaml` 

```text
      ...
      - image:  us.icr.io/<YOUR-NAMESPACE>/os-signal:0.0.1
      ...
```

Update `<YOUR-NAMESPACE>` with your namespace name.

Apply the deployment

```text
kubectl apply -f k8s/private-docker-registry/deployment-private.yaml 
```

Check if the pod got or not.

```text
kubectl get pod
```

```text
private-86d7b54556-4ln47   0/1     ImagePullBackOff   0          3m30s
```

Lets see whats going on by describing the pod

```text
kubectl describe pod -l name=private-pod
```

In the bottom we will see events that happened.

```text
Events:
  Type     Reason     Age                   From                    Message
  ----     ------     ----                  ----                    -------
  Normal   Scheduled  5m2s                  default-scheduler       Successfully assigned default/private-86d7b54556-4ln47 to 10.188.186.13
  Normal   Pulling    3m26s (x4 over 5m1s)  kubelet, 10.188.186.13  pulling image "us.icr.io/mofi-kube/os-signal:0.0.1"
  Warning  Failed     3m26s (x4 over 5m1s)  kubelet, 10.188.186.13  Failed to pull image "us.icr.io/mofi-kube/os-signal:0.0.1": rpc error: code = Unknown desc = failed to resolve image "us.icr.io/mofi-kube/os-signal:0.0.1": no available registry endpoint: pull access denied, repository does not exist or may require authorization: server message: insufficient_scope: authorization failed
  Warning  Failed     3m26s (x4 over 5m1s)  kubelet, 10.188.186.13  Error: ErrImagePull
  Warning  Failed     3m14s (x6 over 5m)    kubelet, 10.188.186.13  Error: ImagePullBackOff
  Normal   BackOff    3m1s (x7 over 5m)     kubelet, 10.188.186.13  Back-off pulling image "us.icr.io/mofi-kube/os-signal:0.0.1"
```

The main line that tells us what went wrong is 

`Failed to pull image "us.icr.io/mofi-kube/os-signal:0.0.1": rpc error: code = Unknown desc = failed to resolve image "us.icr.io/mofi-kube/os-signal:0.0.1": no available registry endpoint: pull access denied, repository does not exist or may require authorization: server message: insufficient_scope: authorization failed`

So our pull access is denied. This is understandable. Our cluster is under the IBM account and the image is in our own account. Kubernetes does not know how to pull this image. 

## Creating Secret for Docker

We will create a docker secret so that Kubernetes can pull this image.

Done Run this yet, but this is what we want.

```text
kubectl create secret docker-registry private-docker-secret \
--docker-server=https://us.icr.io \
--docker-username=iamapikey \
--docker-password=<IAMAPIKEY> \
--docker-email="a@b.com"
```

Well we don't really have that `IAMAPIKEY` at hand. Lets go create that.

First lets log back into your account again.

```text
ibmcloud login --sso
```

Choose your account.

Then run

```text
ibmcloud iam api-key-create docker-iam-key -d "key for accessing docker image registry"
```

This creates a API key. 

> This api key has full user access. Its not good practice to create such api keys. You can also create service api key. Which is a better choice and only gives access to a single service.

The command creates and prints the api key

The key should be stored somewhere as it can't be retrieved later.

```text
Name          docker-iam-key
Description   key for accessing docker image registry
Created At    2019-08-01T06:48+0000
API Key       <Some API Key>
Locked        false
UUID          ApiKey-2e5c6806-ed81-4669-b8cb-e9ba88c832b6
```

Now that we have api key

Lets create the docker secret. But before we need to switch back to the IBM account.   


```text
ibmcloud login --sso
```

Choose the IBM Account

export the apikey for easier copy-paste

```text
export APIKEY="<your-api-key>"
```

```text
kubectl create secret docker-registry private-docker-secret \
--docker-server=https://us.icr.io \
--docker-username=iamapikey \
--docker-password=$APIKEY \
--docker-email="a@b.com"
```

Check the secret was crated

```text
kubectl get secrets
```

We should see our secret there

```text
NAME                    TYPE                                  DATA   AGE
default-au-icr-io       kubernetes.io/dockerconfigjson        1      9h
default-de-icr-io       kubernetes.io/dockerconfigjson        1      9h
default-icr-io          kubernetes.io/dockerconfigjson        1      9h
default-jp-icr-io       kubernetes.io/dockerconfigjson        1      9h
default-token-5tdlr     kubernetes.io/service-account-token   3      9h
default-uk-icr-io       kubernetes.io/dockerconfigjson        1      9h
default-us-icr-io       kubernetes.io/dockerconfigjson        1      9h
kube201-workshop01      Opaque                                2      9h
private-docker-secret   kubernetes.io/dockerconfigjson        1      9m26s
```

Secrets in. Lets try the deployment again.

```text
kubectl get po
```

`imagePullBackOff` again. Why?

So we created a secret but we didn't tell our pod to use it.

In the deployment yaml we have this 2 commented out lines

```text
      containers:
      - image:  us.icr.io/<YOUR-NAMESPACE>/os-signal:0.0.1
        name:  os-signal
        imagePullPolicy: Always
      # imagePullSecrets:
      #   - name: private-docker-secret
```

Lets use nano or vi to uncomment them

```text
nano k8s/private-docker-registry/deployment-private.yaml
```

> imagePullSecrets needs to be at the same level as containers key. Basically delete two characters.

Lets try this one last time. 🤞🏼

```text
 kubectl apply -f k8s/private-docker-registry/deployment-private.yaml
```

After a moment we should get the pod is running.

```text
kubectl get po
```

We know all about what this pod does. But if you want to test its still the same thing. \(You should always test, _**Trust but Verify**_\)

```text
kubectl logs -f -l name=private-pod
```

```text
2019/08/01 07:10:32 awaiting signal
```

If we send a delete command

```text
kubectl delete po -l name=private-pod
```

```text
2019/08/01 07:13:11 Doing all sorts of cleanup work!
2019/08/01 07:13:21 exiting
```

So we can verify the app is working fine. 😎

