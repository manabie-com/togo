const catchAsync = require("../middlewares/async");
const Todo = require("../models/todo");
const User = require("../models/user");
const ApiError = require("../utils/ApiError");

exports.getTodos = catchAsync(async (req, res) => {
    const todos = await Todo.find();
    if (!todos.length) {
        throw new ApiError(404, "No todos");
    }
    res.json({
        sucess: true,
        data: todos
    });
});
exports.createTodo = catchAsync(async (req, res) => {
    const { content } = req.body;
    let todo;
    const userId = req.session.userId;
    const todos = await Todo.findByUserId(userId);
    const user = await User.findById(userId);
    if (todos.length === 0) {
        todo = await Todo.create({
            content,
            userId
        });
    } else {
        const currentDate = new Date();
        let count = 0;
        todos.forEach(todo => {
            if (currentDate.getDate() === todo.createdAt.getDate()) {
                count++;
            }
        });
        if (count === parseInt(user.max_todo)) {
            throw new ApiError(404, "you have added enough to-dos for today");
        } else {
            todo = await Todo.create({
                content,
                userId
            });
        }
    }
    res.status(201).json(todo);
});

exports.deleteTodo = catchAsync(async (req, res) => {
    const { id } = req.params;
    await Todo.findByIdAndDelete(id)
    res.json({
        success: true,
    });
});

exports.updateTodo = catchAsync(async (req, res) => {
    const { id } = req.params;
    const { content } = req.body;
    if (content.length < 3) {
        res.status(404).json({
            sucess: false,
            message: "Must be at least 3 characters"
        });
    } else {
        await Todo.findByIdAndUpdate(id, {
            content
        });
        res.json({
            sucess: true,
        });
    }
});

exports.getTodoById = catchAsync(async (req, res) => {
    const { id } = req.params;
    // console.log(id);
    const getTodoId = await Todo.findById(id);
    if (!getTodoId) {
        throw new ApiError(404, "Not Found");
    }
    res.json({
        success: true,
        data: getTodoId,
    });
});