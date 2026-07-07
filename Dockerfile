ARG ALPINE=alpine:3.22

FROM ${ALPINE} AS builder

RUN apk add --no-cache \
    git \
    gcc \
    musl-dev \
    linux-headers \
    make

WORKDIR /src

# Offizielles Repository klonen
RUN git clone https://github.com/udp-redux/udp-broadcast-relay-redux.git .

# Optional: auf einen bestimmten Release wechseln
# RUN git checkout v1.3.0

RUN make

FROM ${ALPINE}

WORKDIR /runtime

COPY --from=builder /src/udp-broadcast-relay-redux /usr/local/bin/
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]