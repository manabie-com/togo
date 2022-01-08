'use strict!';


const mongoose = require('mongoose');

const taskSchema = new mongoose.Schema({
    userId: mongoose.Schema.Types.ObjectId,
    title: 'string',
    description: 'string',
    dueDate: 'date',
    status: 'number'
}, { timestamps: true });

module.exports.taskModel = mongoose.model('TASK', taskSchema);

