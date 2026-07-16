
docker-entrypoint.sh postgres &

echo "DATABASE_URL=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB" > /app/.env

until pg_isready -q -U postgres; do 
  sleep 1
done

exec /app/contact-management


# docker-entrypoint.sh postgres &

# env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > /app/.env || true

# until pg_isready -q -U postgres; do 
#   sleep 1
# done

# exec /app/contact-management

#!/bin/sh
# docker-entrypoint.sh postgres &

# env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > /app/.env || true

# until pg_isready -q -U postgres; do 
#   sleep 1
# done
# exec /app/contact-management