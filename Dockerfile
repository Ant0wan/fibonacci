FROM golang:1.22.0
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o fibonacci
RUN echo 'fibonacci:x:1000:3000:::' > /etc/passwd

FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
USER fibonacci
COPY --from=0 /app/fibonacci ./
EXPOSE 8000
ENTRYPOINT ["/fibonacci"]
