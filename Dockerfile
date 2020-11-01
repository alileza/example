ARG PROJECT_DIR="/go/src/example"

FROM golang:1.15.2-alpine AS builder

ARG PROJECT_DIR
WORKDIR ${PROJECT_DIR}

COPY . .

RUN apk add make git

RUN make build
RUN mv ${PROJECT_DIR}/bin/example /bin/example

ENTRYPOINT ["example"]