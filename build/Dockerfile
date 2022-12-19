FROM golang AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /go/src/gin-example
COPY . .

RUN go mod vendor
RUN go build -a -o /bin/gin-example ./main.go

FROM alpine AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /bin
COPY --from=builder /bin/gin-example /bin/gin-example

ENTRYPOINT [ "/bin/gin-example" ]