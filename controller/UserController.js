const express       = require('express');
const app           = express();
const UserClass = require('../classes/UserClass');

const user = new UserClass();

module.exports =
{

    async registration(req, res)
    {
        let user_information =
        {
            fullname: req.body.fullname,
            username: req.body.username,
            password: req.body.password,
            confirmpassword: req.body.confirmpassword,
        }

        //console.log(user_information);
        
        let user_class        = new UserClass(user_information);
        let user_valaidation   = await user_class.validate();

        if (user_valaidation.status == "error") {
            res.status(400).send({ message: user_valaidation.message });
        }
        else if(user_valaidation.status == "success")
        {
            let user_registration   = await user_class.create();
            res.send(user_registration);
        }
    },

    async login(req, res)
    {
        let username = req.body.username;
        let password = req.body.password;
        
        let authenticate = await new UserClass({ username: username, password: password }).authenticate();

        if (authenticate.status == "success") {
            res.send(authenticate);
        }
        else {
            res.status(400).send({ message: authenticate.message });
        }
    }
}