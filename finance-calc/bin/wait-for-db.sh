#!/bin/sh
set -e

cmd="$1"

echo "Preparing to check Postgres status..."

until PGUSER=$DB_USER PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -d "$DB_NAME" -c '\q'; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "Postgres is up - executing command"
exec $cmd
