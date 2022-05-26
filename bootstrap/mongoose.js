const mongoose = require('mongoose')

const connect = () => {
    return new Promise(async (resolve, reject) => {
        const dbURI = process.env.NODE_ENV === 'test' ? process.env.MONGODB_URI_TEST : process.env.MONGODB_URI
        try {
            await mongoose.connect(dbURI, {serverSelectionTimeoutMS: 3000})
            console.log('mongodb connected!!!')
        } catch (err) {
            console.log('mongodb error!!!')
            reject(err)
        }
        mongoose.connection.on('error', () => console.log('mongodb error!!!'))
        mongoose.connection.on('disconnected', () => console.log('mongodb disconnected!!!'))
        resolve()
    })
}

module.exports = {mongoose, connect};