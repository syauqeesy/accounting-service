FROM golang:latest

WORKDIR /app
COPY ./user/. .

WORKDIR /common
COPY ./common/. .

WORKDIR /proto
COPY ./proto/. .

RUN apt-get update && apt-get install -y \
  unzip \
  curl \
  git \
  && rm -rf /var/lib/apt/lists/*

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v23.4/protoc-23.4-linux-x86_64.zip \
  && unzip protoc-23.4-linux-x86_64.zip -d /usr/local \
  && rm protoc-23.4-linux-x86_64.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

ENV PATH="/go/bin:${PATH}"

WORKDIR /app
RUN go mod download

WORKDIR /common
RUN go mod download

WORKDIR /proto
RUN go mod download

RUN protoc \
  --proto_path=. \
  --go_out=. \
  --go-grpc_out=. \
  ./*.proto

WORKDIR /
RUN go work init
RUN go work use ./app ./common ./proto

WORKDIR /app
RUN go build -o ./main .
RUN chmod +x start.sh

CMD ["./start.sh"]
