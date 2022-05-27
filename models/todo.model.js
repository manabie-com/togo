const {mongoose} = require('../bootstrap/mongoose')


const ToDoSchema = new mongoose.Schema(
    {
        name: {
            type: mongoose.SchemaTypes.String,
            required: true
        },
        is_completed: {
            type: mongoose.SchemaTypes.Boolean,
            default: false
        },
        user: {
            type: mongoose.SchemaTypes.ObjectId,
            ref: 'user',
            required: true
        },
        created: {
            type: mongoose.SchemaTypes.Date,
            default: new Date()
        },
        updated: {
            type: mongoose.SchemaTypes.Date,
            default: new Date()
        }
    }
)

module.exports = mongoose.model("todo", ToDoSchema)
