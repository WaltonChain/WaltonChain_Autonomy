# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.17-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

ADD . /WaltonChain_Autonomy
RUN cd /WaltonChain_Autonomy && go run build/ci.go install ./cmd/geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /WaltonChain_Autonomy/build/bin/geth /usr/local/bin/

EXPOSE 8645 8646 30503 30503/udp
ENTRYPOINT ["geth"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
