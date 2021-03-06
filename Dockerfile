 FROM golang:1.13-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache gcc musl-dev git

WORKDIR /go/src/monitor

RUN export GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOPRIVATE=*.corp.example.com

# 下载依赖
COPY go.mod go.mod
RUN go mod download

COPY . .

EXPOSE 8001

ENTRYPOINT ["go", "run", "main.go"]
