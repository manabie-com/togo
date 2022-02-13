import mongoose from "mongoose";

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
    task: [
        {
            type: String,
            required: [true, 'Task is required']
        }
    ]
},{
        timestamps: true
});

const User = mongoose.model('User', userSchema)

export default User

