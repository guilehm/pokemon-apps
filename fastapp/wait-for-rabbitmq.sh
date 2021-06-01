#!/bin/bash
set -e

until timeout 1 bash -c "cat < /dev/null > /dev/tcp/spoon-rabbitmq/5672"; do
  >&2 echo "Rabbit MQ not up yet on spoon-rabbitmq"
  sleep 3
done

echo "Rabbit MQ is up"
