FROM golang:1.22-bookworm AS build-env

WORKDIR /src

ENV CGO_ENABLED=0

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -a -o service ./cmd/service 

FROM scratch AS final

COPY --from=build-env /src/service /service

COPY --from=build-env /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/service"]