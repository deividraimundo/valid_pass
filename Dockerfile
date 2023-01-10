#***********BUILDER***********
FROM golang as builder

ARG port_service
ENV PORT_SERVICE $port_service

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./tmp

RUN cd ./tmp && go build -v -o /app/valid_pass

#***********RUNNER***********
FROM ubuntu as runner

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=America/Sao_Paulo

WORKDIR /app

COPY --from=builder /app/valid_pass /app/valid_pass

RUN apt-get update

ENTRYPOINT [ "./valid_pass" ]