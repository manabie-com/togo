const jwt = require('jsonwebtoken');
require('dotenv').config();

// For access token creation
module.exports.createAccessToken = (user) => {
    const data = {
        id: user._id,
        email: user.email
    }

    return jwt.sign(data, process.env.JWT_SECRET, {});
}

// For identity verification
module.exports.verify = (req, res, next) => {
    let token = req.headers.authorization;

    if(typeof token !== 'undefined') {
        token = token.slice(7, token.length);

        return jwt.verify(token, process.env.JWT_SECRET, (err, data) => {
            if(err){
                return res.send({ auth: 'failed' });
            }
            else{
                req.body.userId = data.id;
                next();
            }
        })
    }
    else{
        return res.send({ auth: 'failed' });
    }
}
