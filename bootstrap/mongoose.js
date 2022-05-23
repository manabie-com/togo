const mongoose = require('mongoose')

const connect = () => {
    return new Promise(async (resolve, reject) => {
        if (process.env.NODE_ENV === 'test') {
            const Mockgoose = require('mockgoose').Mockgoose
            const mockgoose = new Mockgoose(mongoose)
            await mockgoose.prepareStorage()
            await mongoose.connect(process.env.MONGODB_URI, {serverSelectionTimeoutMS: 10000})
            console.log('mongodb test mockgoose connected')
            resolve()

        } else {
            try {
                await mongoose.connect(process.env.MONGODB_URI, {serverSelectionTimeoutMS: 3000})
                console.log('mongodb connected!!!')
            } catch (err) {
                console.log('mongodb error!!!')
                reject(err)
            }

            mongoose.connection.on('error', () => console.log('mongodb error!!!'))
            mongoose.connection.on('disconnected', () => console.log('mongodb disconnected!!!'))

            resolve()
        }
    })
}

module.exports = {mongoose, connect};