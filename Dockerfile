FROM alpine:3.7
RUN apk add --no-cache git
ENTRYPOINT ["git"]