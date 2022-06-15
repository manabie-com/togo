var MongoClient = require('mongodb').MongoClient;
var db_name = "togo";
var url = "mongodb://127.0.0.1:27017/"+db_name;

MongoClient.connect(url, function(err, db) {
  if (err) throw err;
  console.log("Database created!");
  db.close();
});

MongoClient.connect(url, function(err, db) {
  if (err) throw err;
  var dbo = db.db(db_name);
  dbo.createCollection("todos", function(err, res) {
    if (err) throw err;
    console.log("todo collection created!");
    db.close();
  });
});


MongoClient.connect(url, function(err, db) {
  if (err) throw err;
  var dbo = db.db(db_name);
  dbo.createCollection("users", function(err, res) {
    if (err) throw err;
    console.log("user collection created!");

    dbo.collection("users").insertOne({user_name: "user01", email: "user01@gmail.com", limit_task_per_day: 1});
    dbo.collection("users").insertOne({user_name: "user02", email: "user02@gmail.com", limit_task_per_day: 2});
    dbo.collection("users").insertOne({user_name: "user03", email: "user03@gmail.com", limit_task_per_day: 3});

    // db.close();
  });
});
