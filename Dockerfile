ARG PROJECT_DIR="/go/src/example"

FROM node:12.18.3-alpine AS builder-fe

COPY ./ui /app

WORKDIR /app

RUN yarn install
RUN yarn build


FROM golang:1.15.2-alpine AS builder

ARG PROJECT_DIR
WORKDIR ${PROJECT_DIR}

COPY . .

RUN apk add make git

COPY --from=builder-fe /app/build /ui/build/

RUN make build
RUN mv ${PROJECT_DIR}/bin/example /bin/example

ENTRYPOINT ["example"]
CMD ["serve"]