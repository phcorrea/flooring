#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 <<-EOSQL
	CREATE DATABASE floor_finder;
	GRANT ALL PRIVILEGES ON DATABASE floor_finder TO postgres;
    CREATE DATABASE floor_finder_test;
	GRANT ALL PRIVILEGES ON DATABASE floor_finder_test TO postgres;
EOSQL