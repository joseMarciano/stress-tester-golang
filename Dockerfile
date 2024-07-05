FROM golang:1.22.5-alpine as builder

COPY . .

RUN go build -o /app/main

FROM scratch
COPY --from=builder  /app/main /app/main
ENTRYPOINT ["/app/main"]