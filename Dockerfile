FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY . /app

RUN go build -o rejection-api

FROM busybox:1.36.1 AS run

COPY --from=build /app/rejection-api /usr/bin/rejection-api
RUN chmod +x /usr/bin/rejection-api

EXPOSE 80

CMD ["/usr/bin/rejection-api"]
