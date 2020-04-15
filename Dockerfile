FROM golang:latest AS builder
ADD . /app/backend
RUN ls /app/backend
WORKDIR /app/backend
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main /app/backend/cmd/main.go

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY ./pkg/migartion_scripts/*init* ./
COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 8081


