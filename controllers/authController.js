const catchAsync = require("../middlewares/async");
const User = require("../models/user");
const ApiError = require("../utils/ApiError");

exports.postUserLogin = catchAsync(async (req, res) => {
    const { username, password } = req.body;
    const found = await User.findByUsername(username);
    // console.log("id:", found._id);
    // console.log("id:", req.session.userId);
    if (found && found.password === password) {
        req.session.userId = found._id;
    } else {
        throw new ApiError(404, "Not Found");
    }
    // console.log("userId1:", req.session.userId);
    res.json({
        success: true,
        message: "Logged in successfully"
    });
});



exports.userLogout = catchAsync(async (req, res) => {
    delete req.session.userId;
    res.json({
        success: true,
        message: "Sign out successful",
    });
});