#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=pa55word psql -h "$host" -U "greenlight" -c '\q'; do
  >&2 echo "Postgres is unavailable"
  >&2 echo PGPASSWORD
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd