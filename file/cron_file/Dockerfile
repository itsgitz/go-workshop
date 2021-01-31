# Get golang image with alpine version
FROM golang:alpine

# Change our directory to /go/src/cron_file
WORKDIR /go/src/cron_file

# Copy all files to image
COPY . .

# Copy the crontab file to crontabs system of image
COPY ./cron/file.cron /etc/crontabs/root

# Build the go program
RUN go build main.go
RUN ls -lah

# Run the crontab system
CMD ["crond", "-f", "-d", "8"]