FROM golang:1.12.9-alpine3.10

RUN apk add --update --no-cache \
      git

WORKDIR /golang_sns_sample
COPY . .

RUN go get github.com/cortesi/modd/cmd/modd
CMD ["modd", "-f", "configs/modd.conf"]


