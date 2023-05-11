
FROM golang:1.19-alpine
WORKDIR /usr/src/app
COPY ./ ./
RUN go mod download
RUN go build -o dist/main
CMD [ "./dist/main" ]