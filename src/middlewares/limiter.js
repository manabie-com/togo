const rateLimit = require("express-rate-limit");
const User = require("../models/userModel")

module.exports.limit = rateLimit({
    windowMs: 24 * 60 * 60 * 1000, // 24 hour window
    max: async (req, res) => {
        const user = await User.findById(req.body.userId)
        
        if(user.isPremium === false){
            return 3
        }else{
            return 5
        }
	},
    message: "Maximum daily limit reached"
});
