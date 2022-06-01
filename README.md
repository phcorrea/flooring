** Dependencies

PostgreSQL with PostGIS extension
Go 1.18

** How to spin up the application

docker-compose build

docker-compose up unit_test

docker-compose up floor_finder

** How to test it

curl localhost:8055/v1/healthcheck

** What needs to be implemented

Partner Creation Handler
Partner Creation Handler Tests
Partner Creation Service
Partner Creation Service Tests

Partner Show Handler
Partner Show Handler Tests
Partner Show Service
Partner Show Service Tests

Partner Search Handler
Partner Search Handler Tests
Partner Search Service
    Partner Query with PostGIS to get all partners inside the operational range
        https://postgis.net/workshops/postgis-intro/geography.html
    Partner Query for materials using array operators (Union &&)
        https://www.postgresql.org/docs/current/functions-array.html#ARRAY-OPERATORS-TABLE
Partner Search Service Tests

CLI command to start the http server