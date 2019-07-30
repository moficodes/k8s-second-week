# Setup

## Get Started

[Click this link to get started](https://cloud.ibm.com/registration?cm_mmc=Email_Events-_-Developer_Innovation-_-WW_WW-_-advocates:roger-osorio,eherrer,mofizur-rahman\title:kubernetesthesecondweek-newyorkcity-7292019\eventid:5d1b7a8709329b07edd42c6f\date:Jul2019\type:workshop\team:global-devadvgrp-newyork\city:newyorkcity\country:unitedstates&cm_mmca1=000019RS&cm_mmca2=10004805&cm_mmca3=M99938765&eventid=5d1b7a8709329b07edd42c6f&cvosrc=email.Events.M99938765&cvo_campaign=000019RS)

## Get Access to Kubernetes Cluster

Once you have you account all set up, it's time to get access to a kubernetes cluster. 

Lab Key: **`kube201`**

[Grant Cluster](https://kube201.mybluemix.net)

![](.gitbook/assets/image%20%281%29.png)

Use the lab key and your IBMid, click agree and submit. This will do some iam magic and connect a 3 node cluster with your account.

## Get Cloud Shell

[Get Access to Cloud Shell](https://cloudshell-pyrk8s-ba.us-south.cf.cloud.ibm.com/)

If it asks for a password use : **`PyRk8sBA`** 

![](.gitbook/assets/screen-shot-2019-07-30-at-10.21.55-am.png)

## Login to IBM Cloud

When using the cloud shell it should already log you to the user you selected in the drop down.

```text
ibmcloud account list
```

This should return 2 results.

If it says not logged in login using `ibmcloud login` and choose the ibm account when prompted.

## Clone Repo

All the example code and yaml files are in github. 

```text
git clone https://github.com/moficodes/k8s-second-week.git
```

```text
cd k8s-second-week
```

Now we are ready to rock and roll.

