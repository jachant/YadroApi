FROM golang:1.23-alpine3.20 AS build

WORKDIR /build

COPY . ./
RUN go mod download  && \
    go build /build/cmd/app/main.go  

FROM alpine:3.21


WORKDIR /app

RUN addgroup -S mygroup && adduser -S -G mygroup myuser
USER myuser

COPY --from=build /build/main ./main

LABEL org.opencontainers.image.authors="atryom" \
      org.opencontainers.image.version="0.1.0" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.description="REST-сервис для работы с погодными данными"

EXPOSE 8000

CMD [ "/app/main" ]
