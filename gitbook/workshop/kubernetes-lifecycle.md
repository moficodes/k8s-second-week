# Kubernetes Lifecycle

This one is a hidden gem. You probably read or heard in the cloud native world we are supposed to gracefully shutdown. What does it mean really? I mean if something bad happened and we had to exit the container what can we really do? Thats where lifecycle hooks are pretty useful.

There are two hooks that are exposed to containers

1. PostStart
2. PreStop

PostStart

This hook executes right after the container is created. It doesn't however guarantee execution before container ENTRYPOINT. This does not take any parameter.

PreStop

This hook is called right before a container is terminated. 

