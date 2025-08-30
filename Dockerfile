# Updated Dockerfile with initialization
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git make build-base linux-headers

WORKDIR /app
COPY . .
RUN make build

FROM alpine:latest

RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app/build/rsuncitychaind /usr/local/bin/rsuncitychaind
COPY init-chain.sh /usr/local/bin/init-chain.sh

RUN chmod +x /usr/local/bin/init-chain.sh

EXPOSE 26656 26657 1317 9090

# Initialize and start
CMD ["sh", "-c", "/usr/local/bin/init-chain.sh && rsuncitychaind start --home /root/.rsuncitychain"]