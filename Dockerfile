FROM alpine:latest
WORKDIR /app
COPY app/app .
COPY app/images ./images
EXPOSE 8080
CMD ["./app"]
