FROM alpine:3.7
RUN apk add --no-cache git && \
    mkdir -p /kubsrv/repo && \
    chgrp -R 0 /kubsrv/repo && \
    chmod -R g+rwX /kubsrv/repo
ENV PORT=3000
ENV CHE_REGISTRY_UPDATE_INTERVAL=60
ENV CHE_REGISTRY_REPOSITORY=/kubsrv/repo
ENV CHE_REGISTRY_GITHUB_URL=https://github.com/skabashnyuk/che-registry.git
EXPOSE 3000
ADD kubsrv_Linux_x86_64 /kubsrv
ENTRYPOINT /kubsrv/kubsrv_Linux_x86_64 -github $CHE_REGISTRY_GITHUB_URL -registry $CHE_REGISTRY_REPOSITORY   -update $CHE_REGISTRY_UPDATE_INTERVAL