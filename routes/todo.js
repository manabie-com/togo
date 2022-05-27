const express = require('express');
const {getAllUsers, getUser, getUserSetting} = require("../controllers/user.controller");
const {addTodo, getAllTodos, checkValidTodoPerDay} = require("../controllers/todo.controller");
const router = express.Router();

/* GET users listing. */
router.get('/', async function (req, res, next) {
    const users = await getAllUsers();
    res.json(users).status(200);
});

router.get('/:userId/todos', async (req, res, next) => {
    try {
        const {userId} = req.params
        await getUser(userId)
        const todos = await getAllTodos(userId)
        return res.json(todos).status(200)
    } catch (err) {
        return res.status(400).json(err)
    }
})

router.post('/:userId/todos', async (req, res, next) => {
    try {
        const {userId} = req.params
        const {name} = req.body
        const [user, userSetting] = await Promise.all([getUser(userId), getUserSetting(userId)])
        await checkValidTodoPerDay(userId, userSetting.limit_per_day)
        const todos = await addTodo(name, userId)
        return res.status(201).send(todos)
    } catch (err) {
        console.log(err)
        if (['invalidUser'].includes(err.errorCode)) {
            return res.status(400).json(err)
        }
        if (err.errorCode === 'limitTodoPerDay') {
            return res.status(406).json(err)
        }
        return res.status(500).json({errorCode: 'server error'})
    }
})

module.exports = router;
