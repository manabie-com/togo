const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const schema = new Schema({
  username: {
    type: String,
    required: true,
    unique: true,
    lowercase: true,
  },
  password: {
    type: String,
    required: true,
  },
  limit: {
    type: Number,
    default: 0,
  },
});

module.exports = mongoose.model("users", schema, "users");
