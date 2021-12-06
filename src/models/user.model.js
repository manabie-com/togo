const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const AccountSchema = new Schema(
  {
    userName: {
      type: String,
      unique: true,
    },  
    password: String,
  },
  { timestamps: true }
);
AccountSchema.index({user: "text"})
module.exports = mongoose.model("Account", AccountSchema);
