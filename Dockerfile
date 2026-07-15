FROM golang:1.26.4-alpine AS build

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o contact-management main.go


FROM postgres:alpine

WORKDIR /app
ENV POSTGRES_USER=contact-management
ENV POSTGRES_PASSWORD=contact-management
COPY --from=build /app/contact-management .
COPY --from=build /app/database/contacts.sql /docker-entrypoint-initdb.d/contacts.sql
COPY --chmod=755 database/entrypoint.sh /app/entrypoint.sh

CMD ["/app/entrypoint.sh"]