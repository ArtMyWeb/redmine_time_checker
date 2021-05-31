FROM golang:1.15-alpine3.12

ENV TZ=Europe/Kiev

WORKDIR /root

COPY . .

RUN apk add --no-cache gcc git libc-dev
RUN go mod tidy
RUN go build -a -ldflags '-extldflags "-static"' .

RUN apk add --no-cache \
		ca-certificates

CMD ["./redmine_time_checker"]
