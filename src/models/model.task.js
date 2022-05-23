'use strict!';


const mongoose = require('mongoose');

const taskSchema = new mongoose.Schema({
    userId: {
        type: mongoose.Schema.Types.ObjectId,
        ref: 'USER'
    },
    title: {
        type: 'string'
    },
    description: {
        type: 'string'
    },
    status: {
        type: 'number'
    }
}, { timestamps: true });

module.exports.taskModel = mongoose.model('TASK', taskSchema);

