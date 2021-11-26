FROM golang:1.17-alpine3.14

WORKDIR / project_backend

COPY . .

RUN go mod download


RUN go build -o mainfile

EXPOSE 8181

CMD ["./mainfile"]