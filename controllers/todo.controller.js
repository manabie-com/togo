const ToDoModel = require('../models/todo.model')
const User = require('../models/user.model')
const {LimitTodoPerDay} = require("../utils/errors");
const moment = require("moment");

const addTodo = async (nameTodo, userId) => {
    const todo = new ToDoModel({
        name: nameTodo,
        user: userId
    })
    await todo.save()
    return todo
}

const checkValidTodoPerDay = async (userId, limitPerDay) => {
    const today = moment().startOf('day')
    const res = await ToDoModel.count({
        user: userId, created: {
            $gte: today.toDate(),
            $lte: moment(today).endOf('day').toDate()
        }
    })
    if (res >= limitPerDay) {
        throw new LimitTodoPerDay(`You have only ${limitPerDay} todos per day`)
    }
}

const getAllTodos = async (userId) => {
    User.find({_id: userId})
    return ToDoModel.find({user: userId}, {}, {});
}

module.exports = {
    addTodo,
    getAllTodos,
    checkValidTodoPerDay
}