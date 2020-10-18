FROM node:alpine as frontend
COPY frontend /frontend/
WORKDIR /frontend
RUN npm i && \
    npm i -g sass && \
    npm run compile

FROM golang:alpine as backend
WORKDIR /go/src/github.com/cosasdepuma/masterchef
COPY main.go go.mod go.sum ./
COPY pkg/ ./pkg/
COPY scripts/front2back.sh ./scripts/
COPY --from=frontend /frontend/dist/index.html ./frontend/dist/index.html
RUN apk update && \
    apk add --virtual essentials --no-cache git upx && \
    sh ./scripts/front2back.sh && \
    GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -a -ldflags="-w -s -extldflags \"-static\"" -o ./bin/masterchef main.go && \
    upx -9 ./bin/masterchef && \
    apk del essentials && \
    rm -rf /var/cache/apk/*

FROM alpine as system
ENV UID 10001
ENV USER appuser
RUN apk update && \
    apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

FROM scratch
COPY --from=system /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=system /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=system /etc/passwd /etc/passwd
COPY --from=system /etc/group /etc/group
COPY --from=backend /go/src/github.com/cosasdepuma/masterchef/bin/masterchef /app/masterchef
USER appuser:appuser
EXPOSE 7767
WORKDIR /app
ENTRYPOINT ["/app/masterchef", "-host", "0.0.0.0"]