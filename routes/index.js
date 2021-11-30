const express = require("express");
const router = express.Router();
const cors = require("cors");
const {checkAuth} = require("../services/auth");
const TOKEN = 'testabc.xyz.ahk';
const fs = require("fs");
const moment = require('moment');
const crypto = require('crypto')

router.use((req, res, next) => {
    res.header("Access-Control-Allow-Origin", req.get("Origin") || "*");
    res.header(
        "Access-Control-Allow-Headers",
        "Origin, X-Requested-With, Content-Type, Accept"
    );
    next();
});
router.use(cors());

router.use(
    express.urlencoded({
        extended: true,
    })
);
router.use(express.json());

router.use(async (req, res, next) => {
    if (req.path === "/login") {
        return next();
    }
    const {authorization} = req.headers;
    if (authorization !== TOKEN) {
        return res.status(401).json({error: "Unauthorized"});
    }
    return next();
});

router.use("/login", async (req, res) => {
    const {username, password} = req.body;

    const isLogin = checkAuth(username, password);
    if (!isLogin) {
        return res.status(401).json({error: "Unauthorized"});
    }
    return res.status(200).json({token: TOKEN, message: "Success"});
});

router.use("/todos", async (req, res) => {
    fs.readFile("./todos.json", "utf8", (err, jsonString) => {
        if (err) {
            return res.status(404).json({error: "Can't find todos.json"});
        }
        try {
            const todos = JSON.parse(jsonString);
            return res.status(200).json({data: todos});
        } catch (err) {
            return res.status(400).json({error: "Can't parse todos.json"});
        }
    });
});

router.use('/createTodo', async (req, res) => {
    const {content, user_id, status} = req.body;
    if (user_id && content && status && (status === 'ACTIVE' || status === 'COMPLETED')) {
        fs.readFile("./todos.json", "utf8", (err, jsonString) => {
            if (err) {
                return res.status(400).json({message: "Content or Status is invalid"});
            }
            try {
                let todos = JSON.parse(jsonString);
                const today = moment()
                const findTodoByUser = todos.filter((todo) => {
                    return (todo.user_id === user_id &&
                        today.diff(moment(todo.created_date), 'days') === 0)
                });

                if (findTodoByUser.length &&
                    findTodoByUser.length > 4) {
                    return res.status(400).json({error: "Limit of 5 tasks " + user_id + " can be added per day."});
                } else {
                    todos.push({
                        "content": content,
                        "created_date": today.toISOString(),
                        "status": status,
                        "id": crypto.createHash('md5').update(today.toISOString()).digest('hex'),
                        "user_id": user_id
                    })
                    fs.writeFile('./todos.json', JSON.stringify(todos), err => {
                        if (err) {
                            return res.status(400).json({error: "Can't add todo into todos.json"});
                        } else {
                            return res.status(200).json({token: TOKEN, message: "Create todo Successfully"});
                        }
                    })
                }
            } catch (err) {
                return res.status(400).json({error: "Can't parse todos.json"});
            }
        });
    } else {
        return res.status(400).json({message: "Content or Status is invalid"});
    }


})
module.exports = router;