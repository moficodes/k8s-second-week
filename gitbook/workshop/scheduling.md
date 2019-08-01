# Scheduling

At the heart of it, Kubernetes is a scheduler. But how does scheduling work? The scheduler will try to put things in the "right" place to the best of its abilities but sometimes we need a bit of extra control. 

## How Does the Scheduler Work?

Kubernetes scheduler is an amazing piece of software. And to explain how it works it will take way too long and I don't think I am even remotely qualified to do the explanation any justice.

Instead You can read [this article](https://medium.com/@dominik.tornow/the-kubernetes-scheduler-cd429abac02f)

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

In kubernetes for most cases we don't care about where the application runs. But there are some performance benefit to having certain pods run close to some other pods. For example if you have 3 nodes in 3 location and you have an application with 3 pods if all those pod run in the same node you basically get no benefit by having a node close to your user. In a case like this you want your pod to be spread out. There is also the case when you have an application that uses some other application heavily \(like a datastore or cache\) having them close to one another helps increase performance.

Lets see an example of this. 

We have 3 nodes and 3 replicas of our pod. We also have 3 replicas of the redis cache.

Just like nodeAffinity pod affinity and anti affinity is also either required or the preferred. If its required it wont schedule until the condition is met. If its preferred, it will schedule but will try first to schedule to the node with highest weight.

Lets look at the affinity rule for the web-store found in `k8s/scheduling/pod-affinity-web-server.yaml`

```text
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - web-store
            topologyKey: "kubernetes.io/hostname"
        podAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - store
            topologyKey: "kubernetes.io/hostname"
```

The anti affinity requires to scheduler to schedule the pod away from any pod with label `values: web-store` which happens to be the label of the web-store pod itself.

It also has affinity towards pods with label `values: store` and won't get scheduled until pods with said label is available. That is the label we are choosing for our redis-cache. That means our pod won't schedule unless there is also a redis-cache available.

Lets see this in action

```text
kubectl apply -f k8s/scheduling/pod-affinity-web-server.yaml
```

If we check with `kubectl get po` we will see the pods are pending.

Lets run the redis cache 

```text
kubectl apply -f k8s/scheduling/pod-affinity-redis-cache.yaml
```

And with that we will see the our web-server pods also get to running.

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

If we want to still schedule something in these nodes we can use something called toleration. A toleration is a way to tell the scheduler that my pod tolerates the taint on this node. 



