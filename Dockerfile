FROM registry.cn-hangzhou.aliyuncs.com/acejilam/golang:1.22.3 AS build

RUN mkdir /src
WORKDIR /src
ENV GOPROXY=https://goproxy.cn/,direct
ADD go.mod go.mod
ADD go.sum go.sum
# RUN go mod download

# ADD pkg pkg
# ADD cmd cmd
COPY . .
# RUN go build  .

FROM alpine
COPY --from=build /src/kelemetry /usr/local/bin/kelemetry

RUN mkdir -p /app/hack
WORKDIR /app
ADD hack/tfconfig.yaml hack/tfconfig.yaml
RUN sed -i 's/127\.0\.0\.1:17272/remote-badger:17271/g' hack/tfconfig.yaml

ENTRYPOINT ["kelemetry"]
