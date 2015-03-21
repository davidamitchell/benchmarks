#!/usr/bin/env ruby

require "json"
require "benchmark"

puts 'hash to json'
3.times do |i|

  n = 1000 ** i
  h = {
    name: "bob jones",
    reference: "his reference"
  }
  beginning_time = Time.now
  time = Benchmark.measure do
    n.times do
      h.to_json
    end
  end

  end_time = Time.now
  elapsed = (end_time - beginning_time) * 1000.0
  each = elapsed / n
  printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", each, n, elapsed)

end

puts 'json to hash'
3.times do |i|
  n = 1000 ** i
  j = "{\"name\":\"bob jones\",\"reference\":\"his reference\"}"
  beginning_time = Time.now
  time = Benchmark.measure do
    n.times do
      JSON.parse(j)
    end
  end

  end_time = Time.now
  elapsed = (end_time - beginning_time) * 1000.0
  each = elapsed / n
  printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", each, n, elapsed)
end
