### Dockerfile (multistage) for wappgo ###

### build stage
FROM golang AS builder

RUN mkdir /wappgo
COPY go.mod go.sum /wappgo/

WORKDIR /wappgo

RUN go mod download

COPY . /wappgo/

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o ./main

### final stage

FROM scratch
WORKDIR /wappgo
COPY --from=builder /wappgo/ /wappgo

EXPOSE 8080

ENTRYPOINT ["./main"]
