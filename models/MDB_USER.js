const MONGOOSE = require('../config/mongo');
const MODEL = require('./MODEL');
const Schema = MONGOOSE.Schema;

const schema = new Schema({
    fullname: {
        type: String,
        require: true
    },
    username: {
        type: String,
        require: true
    },
    password: {
        type: String,
        require: true
    },
    limit: {
        type: Number,
        default: 0
    }
},  {
    timestamps: {
        createdAt: 'created_at',
        updatedAt: 'updated_at',
    }
});

class MDB_USER extends MODEL {
    constructor() {
        super('user', schema);
    }

    async findByUser(user) {
        const res = await this.collection.findOne({ username: user });
        return res ? res.toJSON() : null;
    }
    
    async findOneAndUpdate(doc, updates) {
        const res = await this.collection.findOneAndUpdate(doc, updates, { new: true })
        return res;
    }

}

module.exports = MDB_USER;