COPY . .

RUN go get -tags musl -t ./...

RUN GOOS=linux go build -tags musl -a -installsuffix cgo -ldflags="-w -s" -o app .

FROM alpine:latest

WORKDIR /template_soa

ENV SERVER_PORT=80

COPY --from=builder template_soa /template_soa

RUN apk add git make pkgconf build-base bash

CMD ["./app"]