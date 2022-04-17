const mongoose = require("mongoose");
const app = require("./app");

require("dotenv").config();

const { PORT = 5000 } = process.env;

// Starts the DB then starts the server
mongoose.connect(process.env.MONGO_URL_LOCAL, () => {
  console.log(
    "Successfully connected to the database. Starting the server now..."
  );

  app.server = app.listen(PORT, () => {
    console.log("Server started.");
    console.log("Running on port " + PORT);
  });
});
