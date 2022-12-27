default: image push

image:
	podman build -t dhcp-8-34-226.telco5gran.eng.rdu2.redhat.com:8443/vg/img-pull-go:latest -f Containerfile .
push:
	podman push dhcp-8-34-226.telco5gran.eng.rdu2.redhat.com:8443/vg/img-pull-go:latest
