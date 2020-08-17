# FROM golang:alpine AS builder

# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# WORKDIR /build

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY . .

# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait ./wait
# RUN chmod +x ./wait

# RUN go build -o main .

# WORKDIR /dist

# RUN cp /build/main .
# RUN cp /build/wait .

# FROM scratch


# COPY --from=builder /dist/main /

# EXPOSE 50051

# RUN ls

# CMD ["/wait && ./main"]

FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .


FROM alpine:latest

RUN apk add --no-cache tzdata
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

# Uncomment when there is need to wait for DB connection before start
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait
RUN chmod +x /wait

COPY --from=builder /build/main .

CMD /wait && ./main
