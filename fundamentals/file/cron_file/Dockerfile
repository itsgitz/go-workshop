# Get golang image with alpine version
FROM golang:alpine AS builder

# Change our directory to /go/src/cron_file
WORKDIR /go/src/cron_file

# Copy all files to image
COPY . .

# Build the go program
RUN go build main.go
RUN ls -lah

# Alpine image for shrink the image docker image size
FROM alpine

# Copy go program that has built in builder image to alpine image
WORKDIR /srv
COPY --from=builder /go/src/cron_file .
RUN ls -lah
COPY ./cron/file.cron /etc/crontabs/root

# Run the crontab system
CMD ["crond", "-f", "-d", "8"]