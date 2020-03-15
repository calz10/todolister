FROM golang:latest as builder

LABEL maintainer="Israel Calderon <calderonisrael773@gmail.com>"

ARG buildtime_variable=production

ENV ENV_MODE=$buildtime_variable

WORKDIR /github.com/caldz10/todolister

COPY . .

COPY .env.development .

COPY .env.production .

COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main github.com/calz10/todolister/cmd/todolister

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /root/

ARG buildtime_variable=production

ENV ENV_MODE=$buildtime_variable

COPY --from=builder /github.com/caldz10/todolister/main .

COPY --from=builder /github.com/caldz10/todolister/.env.production .

COPY --from=builder /github.com/caldz10/todolister/.env.development .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 

