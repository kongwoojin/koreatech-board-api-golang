ARG GOARCH=amd64
FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${GOARCH}

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o koreatech-board-api ./cmd
WORKDIR /dist
RUN cp /build/koreatech-board-api .

FROM scratch
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /dist/koreatech-board-api .

EXPOSE 1323

ENTRYPOINT ["/koreatech-board-api"]