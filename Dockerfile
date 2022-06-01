FROM golang:1.18.2-alpine3.16 AS dev

RUN apk update

ENV CGO_ENABLED=0
ENV GOOS=linux

ENV APP_DIR=/usr/local/src/floor_finder/
WORKDIR $APP_DIR

COPY go.* $APP_DIR
RUN go mod download -x

COPY . $APP_DIR
RUN go build -o floor_finder

# Production image
FROM alpine:3.16.0
RUN apk update

RUN adduser -D floorer
USER floorer

ENV APP_DIR=/opt/floor_finder/
WORKDIR $APP_DIR

COPY --from=dev /usr/local/src/floor_finder/floor_finder .

CMD "./floor_finder"
