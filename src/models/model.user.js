'use strict!';


const mongoose = require('mongoose');

const userSchema = new mongoose.Schema({
    username: {
        type: 'string',
        unique: true
    },
    password: 'string',
    limit: 'number'
}, { timestamps: true });

module.exports.userModel = mongoose.model('USER', userSchema);

