# Kubernetes Lifecycle

You probably read or heard in the cloud native world we are supposed to gracefully shutdown. What does it mean really? I mean if something bad happened and we had to exit the container what can we really do? Thats where lifecycle hooks are pretty useful.

## Lifecycle Hooks

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

At this point it will be useful to be able to split our terminal window. 

> _**If you are running this workshop on the cloudshell you can use the tmux shorcut `ctrl+b %` to split the screen. And to move between panes use `ctrl+b <arrow>`**_

On one tab lets run the log to see whats going on

```text
kubectl logs -f -l name=lifecycle-pod
```

This would print the log from the `poststart` state

```text
2019/08/01 05:51:52 file creation has not completed yet...
2019/08/01 05:51:57 file creation has not completed yet...
2019/08/01 05:52:02 file creation has not completed yet...
2019/08/01 05:52:07 file creation has not completed yet...
2019/08/01 05:52:12 file creation has not completed yet...
2019/08/01 05:52:17 file creation has not completed yet...
2019/08/01 05:52:22 file creation has not completed yet...
2019/08/01 05:52:27 file created. starting application.
2019/08/01 05:52:27 starting application in por 8080
```

We also have the `preStop` hook set. We set that to a http call to the path `/shutdown` which we made in our code and simulate some clean up work. 

On the other pane

```text
kubectl delete po -l name=lifecycle-pod
```

We will see in our log the "cleanup work" we simulate in our code. It will take about 10 seconds you will the pod in terminating state while that happens. In our log we should see

```text
...
2019/08/01 05:52:27 file created. starting application.
2019/08/01 05:52:27 starting application in por 8080
2019/08/01 05:53:54 shutdown initiated!
2019/08/01 05:53:54 doing some cleanup work
2019/08/01 05:54:04 done!
```

## OS-Signals

#### Sigterm

When a pod is getting evicted or user calls delete on the pod or the deployment the sigterm signal is sent to the container. We can actually listen for the `sigterm` signal and do things with it. 

Lets deploy the pod with sigterm catch enabled.

```text
kubectl apply -f k8s/lifecycle/deployment-sigterm.yaml
```

The code that corresponds to this version of pod is available at `src/os-signal/main.go` . In the application we basically listen for os signals, and when the sigterm signal comes in we start our "clean up" process. Which right now is just sleeping for 10 seconds and exiting. 

To see this in action, 

Lets first start a tail on the log

```text
kubectl logs -f -l name=signal-pod
```

We should see

```text
2019/08/01 05:56:40 awaiting signal
```

Then run on the other pane

```text
kubectl delete po -l name=signal-pod
```

Your delete command would seemingly hang until the whole delete process ends.

In the log pane we will see the clean up messages. This means we caught the sigterm signal.

```text
2019/08/01 05:58:39 Doing all sorts of cleanup work!
2019/08/01 05:58:49 exiting
```

#### Sigkill

We could also choose to ignore the `sigterm` signal by caching and not exiting. We can see how that is achieved by deploying  

```text
kubectl apply -f k8s/lifecycle/deployment-rouge.yaml
```

Lets see the log

```text
kubectl logs -f -l name=rouge-pod
```

This looks just like our previous pod

```text
2019/08/01 06:01:37 awaiting signal
```

But this pod is rouge. When a delete command is sent it just ignores it.

```
kubectl delete po -l name=rouge-pod
```

On the log you will see the print out. The delete command would just hang there.

```text
2019/08/01 06:01:57 going rouge
2019/08/01 06:01:57 i am invincible
```

But fear not! Kubernetes scheduler has a way to deal with such issue as well. After a sigterm is sent if the pod does not exit in 30 second \(default\) or a time set by the operator sigkill command is sent to the container. This is the `kill -9` for those of us familiar with the linux command. A process has no way to stop and sigkill and the process is terminated. 

