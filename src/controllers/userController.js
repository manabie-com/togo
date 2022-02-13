const User = require('models/userModel');
const auth = require('middlewares/auth');
const bcrypt = require('bcrypt');

// User login
module.exports.login = async (req, res) => {
    const user = await User.findOne({ email: req.body.email });

    if(user === null){
        res.send({ auth: 'user does not exist' });
    }
    else{
        const isPasswordMatched = bcrypt.compareSync(req.body.password, user.password);

        if(isPasswordMatched){
            res.send({ access: authenticator.createAccessToken(user.toObject()) });
        }
        else{
            res.send({ auth: 'login_failed' });
        }
    }
}

// User registration
module.exports.register = async (req, res) => {
    const user = new User({
        name: req.body.name,
        email: req.body.email,
        password: bcrypt.hashSync(req.body.password, 10)
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