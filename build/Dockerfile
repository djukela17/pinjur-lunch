FROM golang:1.11

WORKDIR /app

COPY ./cmd/pinjur-lunch .
COPY ./data ./data
COPY ./web/templates ./web/templates

#run $app_dir/$app_name

EXPOSE 80

ENTRYPOINT ["/app/pinjur-lunch"]
