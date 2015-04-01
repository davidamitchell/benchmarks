var pg = require('pg');

var client = new pg.Client("postgres://localhost:5432/benchmarking");
client.connect();

var hrstart = process.hrtime();

var n = 1000;
for(var i=0; i < n; ++i){
  var query = client.query('select name, reference from bench');

  query.on('end', function(result) {
    console.log(result.rows.length + ' rows were received');
  });

  query.on('error', function(error) {
    console.log(error);
    process.exit(1);
  });
}

//process.exit();
