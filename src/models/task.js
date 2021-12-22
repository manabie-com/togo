const mongoose = require('mongoose')
const validator = require('validator')

const moment = require('moment')

const taskSchema = mongoose.Schema({
    title: {
        type: String,
        required: true
    },
    description: {
        type: String
    },
    createdDate: {
        type: Date,
        required: true,
        default: moment().toDate()
    },
    updatedDate: {
        type: Date,
        default: moment().toDate()
    },
    createdById: {
        type: mongoose.Schema.Types.ObjectId, ref: 'User',
        validate: async (value) => {
            const User = require('./user')
            user = await User.findById(value);
            if (!user) {
                throw new Error('Invalid Created User Id')
            }
        }
    }
})



taskSchema.pre('save', async function (next) {
    // Hash the password before saving the user model
    const task = this;
    if (task.isModified('updatedDate')) {
        task.updatedDate = moment();
    }
    next()
})

const Task = mongoose.model('Task', taskSchema)

module.exports = Task;
