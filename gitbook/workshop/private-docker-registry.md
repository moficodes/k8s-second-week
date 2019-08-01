# Private Docker Registry

Docker Hub Public Images work pretty great. But when we are talking about enterprise software we want to keep the code secure. We also need to keep our images secure, because anyone with access to the docker image essentially has the business logic of your application.

Docker hub is where almost all public images are stored. Which is where we have been getting our images from so far. But for production code its neither secure nor acceptable to put our production code in a public docker registry.

Almost all major cloud provider and some image registry provider including docker gives us a private docker registry. The usage of private docker registry works pretty much the same way. 

We will use ICR or IBM Container Registry that comes built in with your IBM Cloud account.

## Building a Image and Pushing to ICR





