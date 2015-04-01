#!/usr/bin/env ruby

require 'rethinkdb'
require "benchmark"

include RethinkDB::Shortcuts
conn = r.connect(:host=>"localhost", :port=>28015).repl
r.db_create('benchmarking').run(conn) unless r.db_list().run(conn).include?('benchmarking')
conn.close

conn = r.connect(host:"localhost", port: 28015, db: 'benchmarking').repl
r.table_create('bench').run(conn) unless r.table_list().run(conn).include?('bench')

puts "inserting one at a time durablity soft"
3.times do |i|

  n = 2000

  beginning_time = Time.now
  time = Benchmark.measure do
    n.times do
      r.table("bench").insert({name: 'name', reference: 'reference'}).run(durability:"soft")
    end
  end

  end_time = Time.now
  elapsed = (end_time - beginning_time) * 1000.0
  each = elapsed / n
  printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", each, n, elapsed)
end


puts "inserting an array durablity soft"
3.times do |i|

  n = 20000

  beginning_time = Time.now
  time = Benchmark.measure do
    a = []
    n.times do
      a << {name: 'name', reference: 'reference'}
    end
    r.table("bench").insert(a).run(durability:"soft")
  end

  end_time = Time.now
  elapsed = (end_time - beginning_time) * 1000.0
  each = elapsed / n
  printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", each, n, elapsed)
end

puts "inserting one at a time"
3.times do |i|

  n = 200

  beginning_time = Time.now
  time = Benchmark.measure do
    a = []
    n.times do
      r.table("bench").insert({name: 'name', reference: 'reference'}).run()
    end
  end

  end_time = Time.now
  elapsed = (end_time - beginning_time) * 1000.0
  each = elapsed / n
  printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", each, n, elapsed)
end
