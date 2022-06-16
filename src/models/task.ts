import mongoose from 'mongoose'
import { updateIfCurrentPlugin } from 'mongoose-update-if-current'
import { UserDoc } from './user'

interface TaskAttrs {
    description: string,
    title: string,
    userId: string
}

export interface TaskDoc extends mongoose.Document{
    description: string,
    title: string,
    version: number
    userId: string
}

interface TaskModel extends mongoose.Model<TaskDoc>{
    build(attrs: TaskAttrs): TaskDoc
}

const TaskSchema = new mongoose.Schema({
    userId:{
        type: mongoose.Schema.Types.ObjectId,
        ref: 'User',
        required: true
    },
    description: {
        type: String
    },
    title: {
        type: String
    }
},{
    toJSON: {
        transform(doc,ret){
            ret.id = ret._id
            delete ret._id
        }
    },
    timestamps: true
})

TaskSchema.set('versionKey','version')
TaskSchema.plugin(updateIfCurrentPlugin)
TaskSchema.statics.build = (attrs: TaskAttrs) => {
    return new Task(attrs)
}

const Task = mongoose.model<TaskDoc, TaskModel>('Task',TaskSchema)
export { Task}