# GoBuildxAction

Go self-hosted 多架构构建辅助 action

1. Dockerfile

使用 `ARG TARGETOS` 和 `ARG TARGETARCH` 对应 `GOOS` 和 `GOARCH`，动态复制编译结果

```dockerfile
FROM alpine:latest
ARG TARGETOS
ARG TARGETARCH 

RUN apk update && \
    apk upgrade --no-cache && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' >/etc/timezone && \
    rm -rf /var/cache/apk/*

COPY /${TARGETOS}/${TARGETARCH}/runner /usr/bin/runner
RUN chmod +x /usr/bin/runner

WORKDIR /data

ENTRYPOINT [ "/usr/bin/runner" ]

```

2. Action

```yaml
      - name: Get Docker Image Url
        id: image
        env:
          URL: example.docker.registry/repo/name
        run: |
          echo LATEST=${URL}:latest >> $GITHUB_OUTPUT
          echo VERSION=${URL}:${GITHUB_REF/refs\/tags\//} >> $GITHUB_OUTPUT
    
      - name: Gen Go Commands
        id: go_build
        uses: MultiMx/GoBuildxAction@v0.7
        with:
          name: runner
          args: "-gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags '-extldflags \"-static\"'"
          platform: linux/amd64,linux/arm64
          target: ./cmd/search

      - name: Build Binary
        env:
          CGO_ENABLED: 0
        run: ${{ steps.go_build.outputs.commands }}

      - name: Build Docker Image and Push
        id: docker_build
        uses: docker/build-push-action@v2
        with:
          push: true
          context: .
          file: ./Dockerfile
          platforms: linux/arm64,linux/amd64
          tags: ${{ steps.image.outputs.VERSION }},${{ steps.image.outputs.LATEST }}
```