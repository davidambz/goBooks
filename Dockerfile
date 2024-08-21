FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ANABLED=0 go build -ldflags="-w -s" -o server cmd/gobook/main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]