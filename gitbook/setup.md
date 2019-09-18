# Setup

## Get Started

[Click this link to get started](https://ibm.biz/BdzbVY)

> Because Cloud sign up for IBM does not require a credit card. It's kind of hard to stop spam accounts. If too many people try to sign up from the same IP the whitelist will stop it.
>
> Ways to go around it:
>
> 1. Use your phone network \(or phone as hotspot\)
> 2. Use a VPN to change your IP

## Get Access to Kubernetes Cluster

Once you have you account all set up, it's time to get access to a kubernetes cluster. 

**Lab Key**

```text
ikslab
```

**Grant Cluster**

[Click Here](https://devopsdayscmh19.mybluemix.net)

![](.gitbook/assets/image%20%284%29.png)

Use the lab key and your IBMid, click agree and submit. This will do some iam magic and connect a 3 node cluster with your account.

## Get Cloud Shell

[Get Access to Cloud Shell](https://workshop.shell.cloud.ibm.com)

If it asks for a password use : **`ikslab`** 

![](.gitbook/assets/screen-shot-2019-07-30-at-10.21.55-am.png)

> **Select IBM from the dropdown.**

## Login to IBM Cloud

When using the cloud shell it should already log you to the user you selected in the drop down.

```text
ibmcloud account list
```

This should return 2 results.

If it says not logged in login using `ibmcloud login --sso` and choose the ibm account when prompted.

## Get Kubectl Access to Cluster

To see the cluster you have access to run

```text
ibmcloud ks clusters
```

That should show you a cluster .

Now lets setup `kubectl` to work with that cluster.

Run

```text
ibmcloud ks cluster config --cluster <your-cluster-name>
```

This will print out the `export KUBECONFIG=<string>`

Paste that in the terminal. To test that went successfully

```text
kubectl get nodes
```

You should see three nodes. 

## Clone Repo

All the example code and yaml files are in github. 

```text
git clone https://github.com/moficodes/k8s-second-week.git
```

```text
cd k8s-second-week
```

Now we are ready to rock and roll.

