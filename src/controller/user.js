const User = require('../models/user')
const TaskLimit = require('../models/taskLimit')
const moment = require('moment')
const {
    validationResult
} = require("express-validator");

const {
    createUserService,
    loginService
} = require('../services/user')
const createUser = async (req, res) => {
    // Create a new user
    try {
        let result = validationResult(req);
        if (result.errors.length !== 0) { //validator   
            let messages = result.mapped();
            let message = "";

            for (m in messages) {
                message = messages[m].msg;
                break;
            }
            throw new Error(message);
        }
        const resultCreate = await createUserService(req.body)
        res.status(201).send(resultCreate)
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}

const login = async (req, res) => {
    //Login a user
    try {
        let result = validationResult(req);
        if (result.errors.length !== 0) { //validator   
            let messages = result.mapped();
            let message = "";

            for (m in messages) {
                message = messages[m].msg;
                break;
            }
            throw new Error(message);
        }

        const {
            email,
            password
        } = req.body;

        const resultLogin = await loginService(email, password)
        if (!resultLogin) {
            return res.status(401).send({
                error: 'Login failed! Check authentication credentials'
            });
        }
        res.send(resultLogin);
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}

const getUser = async (req, res) => {
    // View logged in user profile
    res.send(req.user)
}

const limitTaskUser = async (req, res) => {
    // Create a new task limit
    try {
        const id = req.params?.id
        if (id) {
            if (!req.body._id) {
                req.body.createdDate = moment().toDate()
            }
            req.body.userId = id
            let taskLimit = new TaskLimit(req.body)
            await TaskLimit.updateOne({
                _id: req.body._id
            }, req.body, {
                upsert: true,
                setDefaultsOnInsert: true
            });
            res.status(201).send({
                taskLimit
            })
        } else {
            res.status(422).send({
                error: "Missing id param"
            });
        }
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
const logOut = async (req, res) => {
    //Log user out of the application
    try {
        req.user.tokens = req.user.tokens.filter((token) => {
            return token.token != req.token
        })
        await req.user.save();
        res.send();
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
const logOutAll = async (req, res) => {
    //Log user out of all devices
    try {
        req.user.tokens = [];
        await req.user.save();
        res.send();
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
module.exports = {
    limitTaskUser,
    createUser,
    login,
    getUser,
    logOut,
    logOutAll
}