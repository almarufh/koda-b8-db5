docker-entrypoint.sh postgres &

env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > /app/.env || true

until pg_isready -q -U postgres; do 
  sleep 1
done

exec /app/contact-management