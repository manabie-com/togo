const {mongoose} = require('../bootstrap/mongoose')


const UserSettingSchema = new mongoose.Schema({
    user: {
        type: mongoose.SchemaTypes.ObjectId, ref: 'user'
    },
    limit_per_day: {
        type: mongoose.SchemaTypes.Number, required: true
    }, created: {
        type: mongoose.SchemaTypes.Date, default: new Date()
    }, updated: {
        type: mongoose.SchemaTypes.Date, default: new Date()
    }
})

module.exports = mongoose.model('user_setting', UserSettingSchema)

