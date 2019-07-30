# Kubernetes Lifecycle

You probably read or heard in the cloud native world we are supposed to gracefully shutdown. What does it mean really? I mean if something bad happened and we had to exit the container what can we really do? Thats where lifecycle hooks are pretty useful.

There are two hooks that are exposed to containers

1. PostStart
2. PreStop

#### PostStart

This hook executes right after the container is created. It doesn't however guarantee execution before container ENTRYPOINT. This does not take any parameter.

#### PreStop

This hook is called right before a container is terminated. 

We have a example app that we can deploy

```text
kubectl apply -f k8s/lifecycle/deployment-lifecycle.yaml
```

We will see that the application is stuck at container creating for about 30 seconds. If you `cat` the yaml file we just ran, 

```text
        ...
        env:
        - name: WAIT_FOR_POST_START
          value: "true"
        imagePullPolicy: Always
        lifecycle:
          postStart:
            exec:
              command:
              - sh
              - -c
              - sleep 30 && echo "Wake up!" > /tmp/poststart
          preStop:
              httpGet:
                port: 8080
                path: shutdown
        ...
```

So for post start we are running a shell script that has a 30 second sleep than it creates a file. In our application we check for the existence of the file before we start our server. You can read the application code at `src/lifecycle/main.go` . 

We also have the `preStop` hook set. We set that to a http call to the path `/shutdown` which we made in our code and simulate some clean up work. It will take about 10 seconds you will the pod in terminating state while that happens.

#### Sigterm

When a pod is getting evicted or user calls delete on the pod or the deployment the sigterm signal is sent to the container. We can actually listen for the `sigterm` signal and do things with it. 



