require('dotenv').config()
const mongoose = require('mongoose')
const URI = "mongodb://mongo:27017/";
const mongoConnect = () => {
    mongoose.connect(URI, {
        useNewUrlParser: true,
        useUnifiedTopology: true,
    }, err => {
        if(err)
            throw err
        console.log('Connect to MongoDB successfully')
    })
}

module.exports = mongoConnect
