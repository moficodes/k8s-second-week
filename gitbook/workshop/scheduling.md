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

To find out nodeIP we can do

```text
kubectl get nodes
```

Once the node is labeled scheduler will automatically schedule the pod on that node.

### Node Affinity

#### Required

Pods will remain pending until a suitable node is found.

#### Preferred

Pods will start even if nothing matches. But if something matches, that node will be given priority. If multiple node matches partially, the node with the highest weight wins.

Lets look at the example we are running.

```text
cat k8s/scheduling/node-affinity.yaml
```

The part we care about is this

```yaml
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                - key: nodetype
                  operator: In
                  values: 
                    - "dev"
                    - "test"
          preferredDuringSchedulingIgnoredDuringExecution:
            - preference:
                matchExpressions:
                  - key: numCores
                    operator: Gt
                    values: 
                      - "3"
              weight: 1
            - preference:
                matchExpressions:
                  - key: location
                    operator: In
                    values: 
                      - "us-east"
                      - "us-south"
              weight: 5
```

Lets break it down.   
We have two kinds of node affinity set. We require the `nodetype` label to be either `dev` or `test` for us to be able to schedule. Our pod will remain pending until this becomes true. 

But the other affinity is preferred and pods will schedule even if they are not met. But they have a weight and if they exists the highest weight will be given priority. 

Lets test it out. 

```text
kubectl apply -f k8s/scheduling/node-affinity.yaml 
```

If you do a `kubectl get po` you will see the pod is in pending.

first let label our nodes with the preferredLabels

```text
kubectl label node <node1> numCores="4"
```

Then label the second node with location us-east

```text
kubectl label node <node2> location="us-east"
```

A quick `kubectl get po` will show its still pending.

Lets add the required label to all nodes.

```text
kubectl label node -l arch=amd64 nodetype=dev
```

And just with that our pod will get scheduled.

Although we put our weight on the preferred it is possible scheduler put it up with any node that would take it at this point. But if we restart the pod we have great chance we will get it to the node we want to.

Also after a pod get scheduled even if we change the label on the node nothing will change. Since label is ignored during execution.

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

