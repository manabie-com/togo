const rateLimit = require("express-rate-limit");
const User = require("../models/userModel")

module.exports.limiter = rateLimit({
    windowMs: 24 * 60 * 60 * 1000, // 24 hour window
    max: async (req, res) => {
        const user = User.findById(req.body.userId)
        
        if(user.isPremium){
            return 10
        }else{
            return 5
        }
	},
    message: "You have reached the maximum number of requests, please try again after 24 hours"
});
