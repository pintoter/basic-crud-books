FROM golang:1.21.1-alpine AS builder

WORKDIR /usr/local/src

# Copy binary
COPY ./.bin/app /usr/local/src/.bin/app

# Copy configs
COPY ./.env /usr/local/src/
COPY ./configs/main.yml /usr/local/src/configs/

RUN apk add --no-cache postgresql-client

CMD ["sh", "-c", "sh ./scripts/db-connection/wait-db.sh && ./.bin/app"]