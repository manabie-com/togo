const express = require("express");
const connectDB = require("./config/db");
const { PORT } = require("./config");
const cookieSession = require('cookie-session')
const authMiddleware = require('./middlewares/auth');
const todoRouters = require("./routers/todoRouters");
const authRouters = require("./routers/authRouters");
const catchError = require("./middlewares/error");
const app = express();

app.use(express.json());
connectDB();
app.use(cookieSession({
    name: 'session',
    keys: [process.env.COOKIE_KEY || 'secret'],
    maxAge: 24 * 60 * 60 * 1000
}));
app.use(authMiddleware);
//router
app.use("/api/v1/auth", authRouters);
app.use("/api/v1/todo", todoRouters);

app.use(catchError);
app.listen(PORT, () => {
    console.log(`Server is runnig on port ${PORT}`);
});
module.exports = app;



