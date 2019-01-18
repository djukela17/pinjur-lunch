FROM golang:1.11

ENV app_name pinjur-lunch-v1
ENV app_dir /app

WORKDIR $app_dir

COPY ./cmd/pinjur-lunch-v1 .
COPY ./data ./data
COPY ./web/templates ./web/templates

EXPOSE 80
