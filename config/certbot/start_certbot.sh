#!/bin/sh
trap 'exit 0' TERM

while :; do
  ./renew_certificates.sh
  sleep 24h & wait $!
done
