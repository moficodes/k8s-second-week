# Resource Quotas

## What it is?

You may have heard, Kubernetes is a container orchestrator. It means K8s manages containers. A big part of managing container is how ever is not containers at all, It's managing the container runtime and resource allocation. 

The  real definition of Kubernetes should say Kubernetes is an api to compute resources with a scheduler. 

In the days of the past you probably would have one application per VM. A VM has some limit based on the hardware it was running on. But in any case, from the point of view of the application it owned the entire world. Figuratively speaking of course.But in a Kubernetes cluster more often than not the application you wrote is not the only thing. More than one application can and will be running in the same cluster. 

So what is a resource quota? It is mechanism by which Kubernetes prevents one pod/app starving other pods by hogging all the resources.

**`Give me X resources in Y increment`**

In Kuberentes declarative model we say, give me x amount of resources in y increment. X is defined as `limits` and y is defined as `requests` . And as of now the things we can set quota now is memory and cpu. Based on the limits and requests set, kubernetes set pod QOS \(Quality of Service\) in 3 categories.

#### Best Effort

When no limit is set so basically x == 0. Since no limit is set scheduler won't know how much space to give the pod at scheduling time. If the cluster is under heavy load and scheduler needs to make room for some other workload these pods are first to get evicted.

Lets deploy some best-effort pods.

If you have cloned the repo already from the root of the repository run

```text
kubectl apply -f k8s/resource-quotas/deployment-best-effort.yaml
```

#### Burstable

Pods are tagged burstable when there is a limit set but the limit is larger than request \(x &gt; y\). This was a little counter intuitive to me at first, since I thought asking for a little at a time was being more empathetic to the api server. But now that I think about it, it makes total sense. With a different limit and request I am making the job of the scheduler harder. Sure it knows there is an upper limit. But it still can't just give me a fix chunk of memory and cpu and start time. When we need more scheduler then has to try to find some more. But it might not be able to. Burstable is the second tier of qos. If there is no more Best Effort pods left and scheduler still needs to make more room, this will go next. 

Run

```text
kubectl apply -f k8s/resource-quotas/deployment-burstable.yaml
```

#### Guaranteed

If the limit and request are equal \(x == y\) the qos is set to be guaranteed. Guaranteed is the best qos we can get. Only a few other types of pods get higher priority than guaranteed. If every pod tell kubernetes how much resource they would need scheduler can block that much space and won't have to keep looking to add more resource. 

Run

```text
kubectl apply -f k8s/resource-quotas/deployment-guaranteed.yaml
```



Once the pods are all running

```bash
kubectl get po -o yaml | grep qos  
```

Basically what we are doing is, for all pod we have we look at the yaml of the current running state and grepping the `qosClass` field.

```yaml
    qosClass: BestEffort
    qosClass: Burstable
    qosClass: Guaranteed
```



