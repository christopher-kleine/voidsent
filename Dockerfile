FROM golang:1.21-alpine AS builder
LABEL org.opencontainers.image.source="https://github.com/Gridanias-Helden/voidsent"

COPY ./ /app

RUN cd /app && \
    CGO_ENABLED=0 go build -ldflags="-s -w" .

# ----------------------------------------

FROM scratch

WORKDIR /app

COPY --from=builder /app /
COPY ./static /static

CMD ["/voidsent"]

EXPOSE 80
