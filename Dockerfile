FROM alpine:3.7
RUN apk add --no-cache git && \
    mkdir -p /kubsrv

ENV PORT=3000
EXPOSE 3000
ADD kubsrv_Linux_x86_64 /kubsrv
ENTRYPOINT ["/kubsrv/kubsrv_Linux_x86_64"]