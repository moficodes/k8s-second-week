# Scheduling

At the heart of it, Kubernetes is a scheduler. But how does scheduling work? The scheduler will try to put things in the "right" place to the best of its abilities but sometimes we need a bit of extra control. 

### Node Name

### Node Selector:

Using labels specified on the node 

To try deploy the deployment with the nodeSelector enabled run

```bash
kubectl apply -f k8s/scheduling/node-selector.yaml
```

You will see the pod remains pending.

This is expected. Lets look at our yaml file

```yaml
    spec:
      containers:
      - image: moficodes/lifecycle:v0.0.1
        name: lifecycle
      nodeSelector:
        gpu: available
      restartPolicy: Always
```

We have a `nodeSelector` of `gpu: available` set. The scheduler will stop the pod from running unless there is a node with the label gpu with value available is present.

We can fix this easily by running

```text
kubectl label node <nodeip> gpu=available
```

Once the node is labeled scheduler will automatically schedule the pod on that node.

### Node Affinity

#### Required

Pods will remain pending until a suitable node is found.

#### Preferred

Pods will start even if nothing matches. But if something matches, that node will be given priority. If multiple node matches partially, the node with the highest weight wins.

### Pod Affinity and Anti Affinity



### Taints and Tolerations

```text
kubectl taint nodes -l arch=amd64 special=true:NoSchedule
```

Notice the `-l arch=amd64` All our nodes has that label. We can technically use any  label. But this one happen to be set on all our nodes. 

Once this taint is set, all our nodes will have the taint set

```text
kubectl describe nodes | grep Taint
```

```bash
Taints:             special=true:NoSchedule
Taints:             special=true:NoSchedule
Taints:             special=true:NoSchedule
```

If we 

