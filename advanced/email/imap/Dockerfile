FROM golang:alpine
LABEL maintainer="Anggit M Ginanjar <anggit.ginanjar@outlook.com>"

WORKDIR /go/src/imap
COPY . .
COPY ./cron/imap.cron /etc/crontabs/root
RUN ls -lah
RUN go mod download
RUN go build main.go
RUN ls -lah

CMD ["crond", "-f", "-d", "8"]
