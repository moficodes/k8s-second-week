# Sidecar

Sidecar is one the most debated pattern in kubernetes. There is argument to be made both for and against this pattern. 

Whether you agree more than one container that shares memory and cpu should be running in the same pod is kind of besides the point. A lot of projects use it and used well it can make some cool stuff happen. 

For example the Istio service mesh uses the envoy as a sidecar and by setting IP table rules intercepts all network traffic to and from the application container. 

There are some specialized version of the sidecar pattern, namely adapter and ambassador pattern.

![](../.gitbook/assets/image%20%283%29.png)

#### Adapter Pattern 

Adapter pattern is useful when you want expose container metrics to an external source but the container does not expose the metric in a format that the external source can understand. 

#### Ambassador Patter

This pattern is used to hide certain implementation complexity from the app by using a proxy address for certain resources like database. Instead of switching connection string in our app using some env variable we can use a container that proxies the request to the right place.

### Example

Lets revisit the init container demo once again. The webpage is loaded at the beginning of the pod. creation and never after. This is not bad, but with a sidecar we can do better.

We will take the same containers but instead of putting the git pull at the init container level we will put it as a sidecar that long polls the git and pulls any latest changes.

Lets deploy

```text
kubectl apply -f k8s/sidecar/deployment.yaml 
```

```text
kubectl apply -f k8s/sidecar/service.yaml 
```

Find the external ip from the service.

Webpage should be there.

But now if we go change our webpage, the change will be picked up in a minute.

### What Happened

Lets look at the containers 

```text
...
      containers:
      - name: app
        image: centos/httpd
        ports:
        - containerPort: 80
          protocol: TCP
        volumeMounts:
        - mountPath: "/var/www/html"
          name: git
        resources:
          limits:
            memory: "100Mi"
            cpu: "300m"
          requests:
            memory: "100Mi"
      - name: poll
        image: axeclbr/git
        volumeMounts:
        - mountPath: /var/lib/data
          name: git
        env:
        - name: GIT_REPO
          value: https://github.com/moficodes/kubernete-second-week-web
        command:
        - "sh"
        - "-c"
        - "git clone $(GIT_REPO) . && watch -n 60 git pull"
        workingDir: /var/lib/data
```

`- "git clone $(GIT_REPO) . && watch -n 60 git pull"` this is the line that does the long polling. 

If you used your own repo you could test it now. But I will change my webpage with something. And we can see what happens.



