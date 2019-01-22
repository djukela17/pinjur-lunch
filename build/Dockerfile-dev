FROM golang:1.11

RUN mkdir /app

WORKDIR /app

COPY ./cmd/pinjur-lunch .
COPY ./data ./data
COPY ./web/templates ./web/templates

EXPOSE 80

ENTRYPOINT ["/app/pinjur-lunch"]
