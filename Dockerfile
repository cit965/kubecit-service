FROM golang:1.21 AS builder

COPY . /src
WORKDIR /src

RUN apt-get update && apt-get install -y --no-install-recommends \
	ca-certificates  \
        netbase \
        unzip \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip && unzip protoc-3.15.8-linux-x86_64.zip -d ./protoc && mv ./protoc/bin/protoc /usr/local/bin/ && rm -rf protoc-3.15.8-linux-x86_64.zip ./protoc

RUN  make init && make all && make build

FROM debian:stable-slim



COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /app/configs

WORKDIR /app



EXPOSE 8000
#EXPOSE 9000
#VOLUME /data/conf

CMD ["./kubecit-service", "-conf", "/app/configs/qa_config.yaml"]
