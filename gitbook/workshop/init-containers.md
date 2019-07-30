# Init Containers

Now we are getting into some advanced kubernetes stuff. Back in lifecycle we talked a bit about `poststart` and `prestop` methods. If I didn't mention at that time, I will say it again. poststart is not a reliable way to make things happen before our cluster runs, because poststart has no guarantee of completion before the actual container running. For times you need that kind of a setup, you want to use something like an init container. Init container run before any other container in the pod and the other containers are only started at init container completion. In that sense init containers are a lot like jobs.Say you have an application that requires a database to be present. You can write an init container that will make sure the database is present and exit on confirmation. Now the rest of your application can start normally.

Lets try to setup a website for our workshop. 

[This Git Repo](https://github.com/moficodes/kubernete-second-week-web) holds the frontend code.

If you want to test it out yourself fork it. 

Lets run the deployment.

```text
kubectl apply -f k8s/init-containers/deployment.yaml
```

That creates the deployment. Let's open this up to the world using a service. We will use a loadbalancer to expose the app.

```text
kubectl apply -f k8s/init-containers/service.yaml
```

To find the IP we can do

```text
kubectl get svc
```

Copy the external IP and got to that ip in your browser.

![](../.gitbook/assets/image%20%282%29.png)

You should see a webpage like this one.

### What Happened

Lets take a deeper look at the deployment.yaml file

```text
...
      initContainers:
      - name: poll
        image: axeclbr/git
        volumeMounts:
        - mountPath: /var/lib/data
          name: git
        command:
        - "git"
        - "clone" 
        - "https://github.com/moficodes/kubernete-second-week-web"
        - "/var/lib/data"
      containers:
      - name: app
        image: centos/httpd
        ports:
        - containerPort: 80
          protocol: TCP
        volumeMounts:
        - mountPath: /var/www/html
          name: git
      volumes:
      - emptyDir: {}
        name: git
...
```

First we have `initContainers` field. In that we run the `axeclbr/git` image. That image has git installed and we can use it to do git pull. We do a git pull and save it to the specified volume mount. From our container which runs httpd server from apache, we also mount the same volume mount at the `/var/www/html` path which httpd serves to the world. We open port 80 for httpd.

Once we expose the deployment with a service we can now access the httpd container at port 80 of the external ip.



