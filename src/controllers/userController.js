const User = require('../models/userModel');
const auth = require('../middlewares/auth');
const bcrypt = require('bcrypt');

// User login
module.exports.login = async (req, res) => {
    const user = await User.findOne({ email: req.body.email });

    if(user === null){
        res.send({ auth: 'User does not exist' });
    }
    else{
        const isPasswordMatched = bcrypt.compareSync(req.body.password, user.password);

        if(isPasswordMatched){
            res.send({ access: auth.createAccessToken(user.toObject()) });
        }
        else{
            res.send({ auth: 'Login failed' });
        }
    }
}

// User registration
module.exports.register = async (req, res) => {
    const user = new User({
        name: req.body.name,
        email: req.body.email,
        password: bcrypt.hashSync(req.body.password, 10),
        isPremium: req.body.isPremium
    });

    user.save().then((user, error) => {
        if(error){
            res.send(error);
        }
        else{
            user.password = undefined;
            res.send(user);
        }
    })
}

// Add task
module.exports.addTask = async (req, res) => {
    const user = await User.findById(req.body.userId);

    user.task.push({ description: req.body.description });

    user.save().then((user, error) => {
        if(error){
            res.send(error)
        }
        else{
            res.send({ isAdded: true});
        }
    })
}

// Get task
module.exports.getTask = async (req, res) => {
    const user = await User.findById(req.body.userId);
    let tasks = [];

    user.task.forEach((allTasks) => {
        if(user){
            tasks.push(allTasks);
        }
    })
    res.send(tasks)
}