FROM migrate/migrate:latest

WORKDIR /migrations

COPY ./*.up.sql .

ENTRYPOINT [ "migrate" ]
