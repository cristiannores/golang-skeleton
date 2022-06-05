# Go image for building the project
FROM golang:alpine as builder

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

RUN mkdir -p $GOPATH/golang-skeleton
WORKDIR $GOPATH/golang-skeleton

COPY . .
RUN go mod vendor
COPY . .

COPY ./shared/utils/config/config.json /
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOPATH/bin/golang-skeleton ./main.go

# Runtime image with scratch container
FROM alpine
ARG VERSION
ENV VERSION_APP=$VERSION

COPY --from=builder /go/bin/ /app/
RUN mkdir -p /shared/utils/config/
COPY --from=builder /config.json /shared/utils/config/
ENTRYPOINT ["/app/golang-skeleton"]
