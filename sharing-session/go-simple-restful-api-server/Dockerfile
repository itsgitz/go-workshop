# multistage build
# alpine
FROM golang:alpine AS builder

WORKDIR /app

# copy all files from host to image
COPY . .

# build 
RUN go build main.go

FROM alpine
COPY --from=builder /app/main /srv

# docker run 
ENTRYPOINT ["/srv/main"]



