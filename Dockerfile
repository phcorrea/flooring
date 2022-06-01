FROM golang:1.18.2-alpine3.16 AS dev

RUN apk update

ENV CGO_ENABLED=0
ENV GOOS=linux

ENV APP_DIR=/usr/local/src/floor_finder/
WORKDIR $APP_DIR
COPY . $APP_DIR

RUN go build

# Production
FROM alpine:3.16.0
RUN apk update

RUN adduser -D floorer
USER floorer

ENV APP_DIR=/opt/floor_finder/
WORKDIR $APP_DIR

COPY --from=dev /usr/local/src/floor_finder/floor* .

CMD "./floor_finder"
