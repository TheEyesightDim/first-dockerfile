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
