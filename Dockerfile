FROM golang:1.22.0 AS builder
#RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o fibonacci

#FROM gcr.io/distroless/base
#
#COPY --from=builder /app/fibonacci-app /fibonacci-app
#
#EXPOSE 8000
#
#ENTRYPOINT ["/fibonacci-app"]

