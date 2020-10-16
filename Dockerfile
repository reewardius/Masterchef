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
    go build -ldflags "-s -w" -i -o ./bin/masterchef main.go && \ 
    apk del essentials && \
    rm -rf /var/cache/apk/*

FROM scratch
COPY --from=backend /go/src/github.com/cosasdepuma/masterchef/bin/masterchef /masterchef
ENTRYPOINT ["/masterchef"]
