# build frontend static files
FROM node:12.18.3-alpine AS builder-ts

COPY . /app

WORKDIR /app/ui

RUN yarn install
RUN yarn build

# build binary
FROM golang:1.15.2-alpine AS builder-go
RUN apk add make git

WORKDIR /go/src/example

COPY . .
RUN make build


# final image
FROM alpine:3.12.1

COPY --from=builder-ts /app/ui/build /ui/build/
COPY --from=builder-go /go/src/example/bin/example /bin/example
COPY --from=builder-go /go/src/example/autogen/docs /docs

ENTRYPOINT ["example"]
CMD ["serve", "--swagger-path=/docs/example.swagger.json"]