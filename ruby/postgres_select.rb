#!/usr/bin/env ruby

require 'pg'
require "benchmark"


# Output a table of current connections to the DB
3.times do |i|
  conn = PG.connect( dbname: 'benchmarking' )

  n = 20000

  beginning_time = Time.now
  time = Benchmark.measure do
    n.times do
      conn.exec( "SELECT name, reference from bench" ) do |result|
        result.each do |row|
          name      = row['name']
          reference = row['reference']
          # puts reference, name
        end
      end
    end
  end

  end_time = Time.now
  elapsed = (end_time - beginning_time) * 1000.0
  each = elapsed / n
  printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", each, n, elapsed)
end
