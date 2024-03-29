# Daemonsets

Daemon is a process in linux that runs in the background and does certain things. For example if you have ever used docker, the docker cli tool speaks to the docker daemon running in the background that understands our cli command and does the things we want to have happen.

In kubernetes Daemonsets are a special set of pods that run on every schedulable node. 

Why would we want daemon set. If we remember correctly, we achieved similar outcome using pod `affinity` and `antiaffinity`. How is Daemonsets any different. The main difference is Daemonset runs one pod per node. So if the number of pods change daemonset pods change accordingly. In a replicaset we would have to manually intervene to achieve the same result.

Daemonsets are particularly useful when we have to run something at the host level. Like log collecting or device metrics.

If we run

```text
kubectl apply -f k8s/daemonsets/daemonset.yaml
```

We will see three fluentd pod is created.

If we describe each of these pods we can also confirm they were created in different nodes.

