const mongoose = require("mongoose");

const connectDB = () => {
    mongoose.connect(process.env.MONGODB_URL || 'mongodb+srv://pmchauuu:minhchau2510@cluster0.hgaoz.mongodb.net/mytodos?retryWrites=true&w=majority').then(() => {
        console.log("Connect to DB successfully!");
    }).catch(err => console.log("can not connect toDB"));
}

module.exports = connectDB