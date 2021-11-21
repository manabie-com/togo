const mongoose = require("mongoose");
const Schema = mongoose.Schema;
const User = require('./user');
const todoSchma = Schema({
    content: {
        type: String,
        required: [true, "Content is required"],
        minlenght: [3, "Must be at least 3 characters"],
        maxlenght: [30, "must have at most 30 characters"],
    },
    userId: {
        type: mongoose.Schema.Types.ObjectId,
        ref: 'User'
    }
},
    {
        timestamps: true,
    }
);
const Todo = mongoose.model("todos", todoSchma);
Todo.findByUserId = async function (userId) {
    return Todo.find({
        where: {
            userId,
        },
    });
};
module.exports = Todo;