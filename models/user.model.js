const {mongoose} = require('../bootstrap/mongoose')


const UserSchema = new mongoose.Schema(
    {
        username: {
            type: mongoose.SchemaTypes.String,
            required: true
        },
        password: {
            type: mongoose.SchemaTypes.String,
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

module.exports = mongoose.model("user", UserSchema)