# Start by building the application.
FROM golang:1.12 as build

# WORKDIR /go/src/app
# COPY . .

COPY . $GOPATH/src/github.com/CelsoSantos/knative-gloo-grpc/
WORKDIR $GOPATH/src/github.com/CelsoSantos/knative-gloo-grpc/

# RUN go-wrapper download
# RUN go-wrapper install
# RUN go get -d -v ./...
# RUN go install -v ./...

RUN go get -d -v
RUN go build -o /go/bin/app
# RUN go build -o $GOPATH/src/github.com/CelsoSantos/knative-gloo-grpc/

# RUN go mod download
# RUN go build -o /app main.go

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /

CMD ["/app"]

EXPOSE 80