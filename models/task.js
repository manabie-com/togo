var mongoose = require('./../store/mongo_store');
var Schema = mongoose.Schema;
var autoIncrement = require('mongoose-auto-increment');

autoIncrement.initialize(mongoose.togo);

/**
 * Task Schema
 */
const taskSchema = new Schema({
    task_id: {
        type: Number,
        required: true
    },
    task_data: {
        type: String,
        required: true
    },
    task_date: {
        type: Date,
        default: Date.now
    },
    updated: {
        type: [Date]
    },
    user_id: {
        type: Number,
        required: true
    }
}, {collection: 'Task'});

postSchema.plugin(autoIncrement.plugin, {
    model: 'Task',
    field: 'task_id',
    startAt: 1
});

postSchema.index({
    task_data: 'text',
  }, {
    weights: {
        task_data: 1,
    },
  });

let Task = mongoose.togo.model('Task', taskSchema);

module.exports = Task;
