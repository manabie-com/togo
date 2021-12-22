const mongoose = require('mongoose')


const taskLimitSchema = mongoose.Schema({
    quantity: {
        type: mongoose.Schema.Types.Number,
        require: true
    },
    atDate: {
        type: Date,
        required: true
    },
    createdDate: {
        type: Date,
        required: true
    },
    userId: {
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
const TaskLimit = mongoose.model('TaskLimit', taskLimitSchema)

module.exports = TaskLimit;
