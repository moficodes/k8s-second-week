# Scheduling

At the heart of it, Kubernetes is a scheduler. But how does scheduling work? The scheduler will try to put things in the "right" place to the best of its abilities but sometimes we need a bit of extra control. 

### Node Name

### Node Selector:

Using labels specified on the node 

### Node Affinity

#### Required

Pods will remain pending until a suitable node is found.

#### Preferred

Pods will start even if nothing matches. But if something matches, that node will be given priority. If multiple node matches partially, the node with the highest weight wins.

### Pod Affinity and Anti Affinity

### Taints and Tolerations



