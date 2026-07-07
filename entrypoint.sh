#!/bin/sh

set -e

set --

if [ -n "$RELAY_ID" ]; then
    set -- "$@" --id "$RELAY_ID"
fi

if [ -n "$PORT" ]; then
    set -- "$@" --port "$PORT"
fi

OLDIFS=$IFS
IFS=,

for iface in $INTERFACES
do
    if [ -n "$iface" ]; then
        set -- "$@" --dev "$iface"
    fi
done

IFS=$OLDIFS

if [ -n "$MULTICAST" ]; then
    set -- "$@" --multicast "$MULTICAST"
fi

echo
echo "========================================="
echo "Starting udp-broadcast-relay-redux"

echo
echo "Binary:"
./udp-broadcast-relay-redux --help 2>/dev/null | head -1 || true

echo
echo "Interfaces:"
for iface in $(echo "$INTERFACES" | tr ',' ' ')
do
    echo "  - $iface"
done

echo
echo "Arguments:"
echo "  $*"

echo "========================================="
echo

for iface in $(echo "$INTERFACES" | tr ',' ' ')
do
    if ! ip link show "$iface" >/dev/null 2>&1; then
        echo "ERROR: Interface '$iface' not found!"
        exit 1
    fi
done

exec ./udp-broadcast-relay-redux "$@"