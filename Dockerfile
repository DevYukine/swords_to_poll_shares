# syntax=docker/dockerfile:1.7-labs
############################################
# Builder
############################################
FROM --platform=$BUILDPLATFORM golang:1.25-bookworm AS builder

# ---- Build-time configuration (override with --build-arg) ----
ARG VERSION="v0.0.1"
ARG BINARY="swords_to_poll_shares"
ARG CMD_PATH="./cmd/swords_to_poll_shares"

# ---- Environment for reproducible, static builds ----
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct

WORKDIR /src

# 1) Prime the module cache (separate from source for better caching)
COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download


COPY . .


RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -tags netgo \
      -trimpath -buildvcs=false \
      -ldflags "-s -w -X main.version=${VERSION}" \
      -o /out/${BINARY} ${CMD_PATH}

########################################################
# Final image
########################################################
FROM gcr.io/distroless/static-debian12:nonroot AS runtime

ARG VERSION
ARG BINARY

WORKDIR /app
COPY --from=builder /out/${BINARY} /usr/local/bin/${BINARY}

ENV TZ=UTC
EXPOSE 50051

USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/${BINARY}"]
