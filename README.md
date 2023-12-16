# Introduction to Making Container Images With Dockerfiles

This example creates a container image for a Go http server that returns a simple response to requests on an endpont.

## How to Build and Push This Image to a Registry

This image was created with Podman in mind, but the process should be similar in Docker or other container managers.

To build, run:

`podman build -t <[username/]image-name>`

where `username` is your username on a container registry. If you are just building an image for local use, then omit the username.

Log into the registry with:

`podman login <registry_name>`

and enter any necessary authentication info.

Finally, push it to the registry with 

`podman push <username/image-name[:tag-name]>`

assuming you built the image with your username in the first place.

You should then be able to pull the image from the registry with 

`podman pull <registry/username/image-name[:tag-name]>

## Generate a Pod and Add Containers

In Podman, a 'Pod' is a group of containers sharing storage and network resources, with a specification for how to run the containers together. (as per (https://kubernetes.io/docs/concepts/workloads/pods/)[Kubernetes docs.])

To create a pod, run:

`podman pod create`

This creates a pod with a base 'infra' container which allows other containers to be added, started, and stopped inside the pod.
You can also assign ports to this pod on creation by appending `-p host-port:pod-port`.

To add a container to the existing pod, run:

`podman run -dt --pod <podname> <container-name>`

Now, to create a kubernetes YAML manifest file based on this pod configuration, run:

`podman generate kube <podname> >> <filename>.yaml`

Review the contents of the YAML file and see how it describes the state of the pod.

Now, to run back or 'play' that configuration, run:

```
podman pod stop <podname>
podman pod rm <podname>
podman play kube <filename>.yaml
```

This will recreate the pod with its constituent containers, volumes, and configuration, and run it.

