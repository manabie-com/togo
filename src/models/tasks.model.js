const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const schema = new Schema({
  name: {
    type: String,
    required: true,
    unique: true,
  },
  user: {
    type: mongoose.Schema.Types.ObjectId,
    ref: "users",
  }
}, {
  timestamps: true,
});

module.exports = mongoose.model("tasks", schema, "tasks");
