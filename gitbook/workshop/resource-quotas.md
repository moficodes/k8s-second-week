# Resource Quotas

## What it is?

You may have heard, Kubernetes is a container orchestrator. It means K8s manages containers. A big part of managing container is how ever is not containers at all, It's managing the container runtime and resource allocation. 

The  real definition of Kubernetes will say Kubernetes is an api to compute resources with a scheduler. 

In the days of the past you probably would have one application per VM. A VM has some limit based on the hardware it was running on. But in any case, from the point of view of the application it owned the entire world. Figuratively speaking of course.But in a Kubernetes cluster more often than not the application you wrote is not the only thing. More than one application can and will be running in the same cluster. 

So what is a resource quota? It is mechanism by which Kubernetes prevents one pod/app starving other pods by hogging all the resources.







