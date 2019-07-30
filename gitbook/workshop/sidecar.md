# Sidecar

Sidecar is one the most debated pattern in kubernetes. There is argument to be made both for and against this pattern. 

Whether you agree more than one container that shares memory and cpu should be running in the same pod is kind of besides the point. A lot of projects use it and used well it can make some cool stuff happen. 

For example the Istio service mesh uses the envoy as a sidecar and by setting IP table rules intercepts all network traffic to and from the application container. 



