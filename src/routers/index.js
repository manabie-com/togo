const authRoute = require('./user');
const taskRoute = require('./task');

const route=(app)=>{
    app.use('/user',authRoute);
    app.use('/task', taskRoute);
}
module.exports = route;