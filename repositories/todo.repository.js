const Todo = require('../models/todo.model');
const User = require('../models/user.model');
var {validationResult} = require('express-validator');

exports.postTodoCreate = (req, res, next) => {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
        res.status(422).json({ errors: errors.array() });
        return;
      }

    const data = req.body;
    User.findOne({ _id: (data.user_id).trim() }).then(user => {
        console.log(user);
        if(user != null){
            var start = new Date();
            start.setUTCHours(0,0,0,0);
            var end = new Date();
            end.setUTCHours(23,59,59,999);
            Todo.find({created_at: {$gte: start, $lt: end}}).then(tdo => {
                // check user limited todo per day
                if(tdo.length <= user.limit_task_per_day){
                    let todo = new Todo({
                        title: data.title,
                        description: data.description,
                        user_id: data.user_id,
                        created_at: new Date()
            
                    })
                    todo.save()
                        .then(tdo => {
                            console.log(tdo)
                        })
                        .catch(err => {
                            res.json(err);
                        })
                    res.json(tdo);
                }
                res.json({errors: [
                    {
                        "msg": "Limited todo per day, maximum is " + user.limit_task_per_day 
                    }
                ]});
            }).catch(err => {
                console.error(err)
            });
        }else{
            res.json({errors: [
                {
                    "msg": "User does not exist"
                }
            ]});
        }
    })
        .catch(err => {
            console.error(err)
        })
}

exports.getTodos = (req, res, next) => {
    Todo.find({}).then(tdos => {
        res.json(tdos);
    }).catch(err => {
        console.error(err)
    });
}

