require('./environment')
const mongoose = require('mongoose')
const dbUrl = `mongodb://localhost:27017/${process.env.DB_NAME}`
let conn
const connectDB = () => {
    try {
        conn = mongoose.connect(dbUrl, {
            useNewUrlParser: true,
            useUnifiedTopology: true
        })
        return mongoose.connection
    } catch (err) {
        process.exit(1)
    }
}


const getDBConn = () => {
    return conn;
}
module.exports = {
    connectDB,
    getDBConn
}
