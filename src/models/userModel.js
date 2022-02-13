const mongoose = require('mongoose');
const uniqueValidator = require('mongoose-unique-validator');

const schema = new mongoose.Schema({
    name:{
        type: String,
        required: [true, 'Name is required']
    },
    email:{
        type: String,
        unique: true,
        required: [true, 'Email is required']
    },
    password:{
        type: String,
        required: [true, 'Password is required']
    },
    task: [
        {
            type: String,
            required: [true, 'Task is required']
        }
    ]
},{
        timestamps: true
});

schema.plugin(uniqueValidator);
module.exports = mongoose.model('User', schema);

