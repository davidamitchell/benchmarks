#!/usr/bin/env ruby
require 'pg'
require "benchmark"

conn = PG.connect( dbname: 'benchmarking' )
conn.exec( 'drop table if exists bench; create table bench ( name text, reference text );' )
conn.prepare("insert", "insert into bench (name, reference) values ($1, $2)")

3.times do |i|

  n = 20000


  beginning_time = Time.now
  time = Benchmark.measure do
    n.times do
      conn.exec_prepared("insert", ['name', 'reference'])
    end
  end

  end_time = Time.now
  elapsed = (end_time - beginning_time) * 1000.0
  each = elapsed / n
  printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", each, n, elapsed)
end
conn.close
