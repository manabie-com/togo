const mongoose = require('mongoose')
const validator = require('validator')
const bcrypt = require('bcryptjs')
const jwt = require('jsonwebtoken')
const Task = require('./task')
const TaskLimit = require('./taskLimit')
const moment = require('moment')

const userSchema = mongoose.Schema({
    name: {
        type: String,
        required: true,
        trim: true
    },
    email: {
        type: String,
        required: true,
        unique: true,
        lowercase: true,
        validate: value => {
            if (!validator.isEmail(value)) {
                throw new Error({error: 'Invalid Email address'})
            }
        }
    },
    password: {
        type: String,
        required: true,
        minLength: 7
    },
    tokens: [{
        token: {
            type: String,
            required: true
        }
    }]
})

userSchema.pre('save', async function (next) {
    // Hash the password before saving the user model
    const user = this
    if (user.isModified('password')) {
        user.password = await bcrypt.hash(user.password, 10)
    }
    next()
})

userSchema.methods.generateAuthToken = async function() {
    // Generate an auth token for the user
    try {
        const user = this
        const token = jwt.sign({_id: user._id}, process.env.JWT_KEY||"NhanPhan")
        user.tokens = user.tokens.concat({token})
        const result = await User.findOneAndUpdate({_id: user._id}, {
            tokens: user.tokens
        }, {new: true});
        if(result) return token;
        else throw Error('Task not exist')
    } catch (error) {
        throw error
    }
    
}

userSchema.statics.findByCredentials = async (email, password) => {
    // Search for a user by email and password.
    try {
        const user = await User.findOne({ email} )
        if (!user) {
            throw new Error('User is not exist')
        }
        const isPasswordMatch = await bcrypt.compare(password, user.password)
        if (!isPasswordMatch) {
            throw new Error('Password is wrong')
        }
        return user
    } catch (error) {
        throw error
    }
    
}

userSchema.methods.checkTaskPerDayLimit = async (id) => {
    try {
        const today = moment().startOf("day");
        const dateQuery = {
            $gte: today.toDate(),
            $lte: moment(today).endOf('day').toDate(),
        }
        const taskLimit = await TaskLimit.findOne({
            userId: id,
            atDate: dateQuery,
        })
        const limitQuantity = taskLimit?taskLimit.quantity:5;
        let countTask = await Task.count({
            createdById: id,
            createdDate: dateQuery,
        });
        
        if (countTask >= limitQuantity) return true;
        return false;
    } catch (error) {
        throw error;
    }
    
}

const User = mongoose.model('User', userSchema)

module.exports = User
