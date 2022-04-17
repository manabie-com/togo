const mongoose = require("mongoose");

const Schema = mongoose.Schema;

const userSchema = new Schema({
  user_name: { type: String, required: true },
  task_daily_limit: { type: Number, required: true },
});

module.exports = mongoose.model("User", userSchema);
