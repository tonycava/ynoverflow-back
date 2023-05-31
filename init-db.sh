#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER docker;
	DROP DATABASE IF EXISTS ynoverflow;
	CREATE DATABASE ynoverflow;
	GRANT ALL PRIVILEGES ON DATABASE ynoverflow TO docker;
EOSQL