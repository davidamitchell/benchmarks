# Benchmarking

## Languages
* ruby 2.1.5
* golang 1.4
* nodejs 0.12.0

## Databases
* postgres 9.4.1
* rethinkdb
* couchbase

## What
* simple inserts
* simple reads

## How

psql -h localhost < sql/create_database.sql
psql -h localhost -d benchmarking < sql/create_tables.sql

### Json
ruby ruby/json_parse.rb
go run go/json_parse.go

### Database selects
ruby ruby/postgres_select.rb
go run go/postgres_select.go



dropdb benchmarking
