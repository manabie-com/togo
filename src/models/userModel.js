const mongoose = require('mongoose');

const userSchema = new mongoose.Schema({
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
    isPremium:{
        type: Boolean,
        default: false,
    },
    task:[
        {
            description: {
                type: String,
                required: [true, 'Description is required']
            },
            added:{
                type: Date,
                default: new Date()
            }
        }
    ]
},{
        timestamps: true
});

module.exports = mongoose.model('User', userSchema);

