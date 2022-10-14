FROM mysql:8.0.23
COPY ./.env /docker-entrypoint-initdb.d/
COPY ./database/*.sql /docker-entrypoint-initdb.d/