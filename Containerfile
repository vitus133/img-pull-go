FROM golang:1.19 AS builder
WORKDIR /go/src/github.com/vitus133/img-pull-go
COPY . .
RUN go build -mod vendor -buildmode=exe -tags "exclude_graphdriver_devicemapper containers_image_openpgp exclude_graphdriver_btrfs" -o bin/img-pull-go


FROM registry.access.redhat.com/ubi8/ubi:8.7-1037
RUN dnf -y module enable container-tools:rhel8; dnf -y update; rpm --restore --quiet shadow-utils; \
dnf -y install fuse-overlayfs /etc/containers/storage.conf --exclude container-selinux; \
rm -rf /var/cache /var/log/dnf* /var/log/yum.*
ADD containers.conf  policy.json storage.conf /etc/containers/
RUN mkdir -p /var/lib/shared/overlay-images /var/lib/shared/overlay-layers; \
touch /var/lib/shared/overlay-images/images.lock; \
touch /var/lib/shared/overlay-layers/layers.lock


COPY --from=builder /go/src/github.com/vitus133/img-pull-go/bin/img-pull-go /usr/local/bin/

