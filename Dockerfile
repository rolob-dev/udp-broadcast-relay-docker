ARG ALPINE=alpine:3.22

#
# Builder
#
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

# Kompilieren und Build-Ausgabe anzeigen
RUN make && \
    echo "=========================================" && \
    echo "Build directory:" && \
    pwd && \
    echo && \
    echo "Files:" && \
    ls -lah && \
    echo "========================================="

#
# Runtime
#
FROM ${ALPINE}

WORKDIR /runtime

# Binary übernehmen
COPY --from=builder /src/udp-broadcast-relay-redux /usr/local/bin/udp-broadcast-relay-redux

# Entrypoint übernehmen
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]