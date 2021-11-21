const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const userSchma = Schema({
    username: {
        type: String,
        required: [true, "username is required"],
        minlenght: [3, "Must be at least 3 characters"],
        maxlenght: [30, "must have at most 30 characters"],
    },
    password: {
        type: String,
        required: [true, "Password is required"],
        minlenght: [3, "Must be at least 3 characters"],
        maxlenght: [30, "must have at most 30 characters"],
    },
    max_todo: {
        type: String,
    },

}, {
    timestamps: true,
});
const User = mongoose.model("users", userSchma);
User.findByUsername = async function (username) {
    return User.findOne({
        where: {
            username,
        },
    });
};
module.exports = User;