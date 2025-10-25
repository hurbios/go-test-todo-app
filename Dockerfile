FROM golang:tip-trixie
ENV PORT=80
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/server

EXPOSE 80

CMD ["app"]