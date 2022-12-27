# Image puller
This is a test program aimed to study image pulling within a kubernetes pod using go containers libraries.
It is inspired by [this blog post](https://iximiuz.com/en/posts/working-with-container-images-in-go/).
I added a [Containerfile](Containerfile) for building the image and adding everything needed to run the program within a container.

## Usage
### Building and pushing
Adjust the [Makefile](Makefile) with your registry details.
Then 
```bash
$ make
```
### Run the pod and check the log
```bash
$ oc apply -f po.yaml

pod/image-puller created

$ oc get po

NAME           READY    STATUS      RESTARTS   AGE
image-puller   0/1      Completed   0          8s

$ oc logs image-puller

Getting image source signatures
Copying blob sha256:2123501b93d459033750d3ea725953060ed9bb83bac7c13e46c675be22b69f4a
Copying config sha256:827365c7baf137228e94bcfc6c47938b4ffde26c68c32bf3d3a7762cd04056a5
Writing manifest to image destination
Storing signatures
Image manifest:
{
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
   "config": {
      "mediaType": "application/vnd.docker.container.image.v1+json",
      "size": 1457,
      "digest": "sha256:827365c7baf137228e94bcfc6c47938b4ffde26c68c32bf3d3a7762cd04056a5"
   },
   "layers": [
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 2587828,
         "digest": "sha256:2123501b93d459033750d3ea725953060ed9bb83bac7c13e46c675be22b69f4a"
      }
   ]
}
```

## Conclusions
Although the minimum program for pulling an image is simple, the image size is substantial. Compared with the official podman size below.
```bash
$ podman image ls

REPOSITORY                                                        TAG         IMAGE ID      CREATED            SIZE
dhcp-8-34-226.telco5gran.eng.rdu2.redhat.com:8443/vg/img-pull-go  latest      86b7f491dff0  10 minutes ago     275 MB
registry.access.redhat.com/ubi8/podman                            8.7-5       3e7d030126c5  7 weeks ago        361 MB
```
