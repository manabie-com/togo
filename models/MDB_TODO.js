const MONGOOSE = require('../config/mongo');
const MODEL = require('./MODEL');
const Schema = MONGOOSE.Schema;

const schema = new Schema({
    task:
    {
        type: String,
        required: false,
    },
    userName:
    {
        type: String,
        required: false,
    }
},  {
    timestamps: {
        createdAt: 'created_at',
        updatedAt: 'updated_at',
    }
});

class MDB_TODO extends MODEL {
    constructor() {
        super('todo', schema);
    }


}

module.exports = MDB_TODO;