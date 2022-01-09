'use strict!';


const mongoose = require('mongoose');

const userSchema = new mongoose.Schema({
    username: {
        type: 'string',
        unique: true
    },
    password: {
        type: 'string'
    },
    limit: {
        type: 'number',
        default: 10
    }
}, { timestamps: true });

module.exports.userModel = mongoose.model('USER', userSchema);

