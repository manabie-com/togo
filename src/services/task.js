const Task = require('../models/task')
const User = require('../models/user')
const moment = require('moment')


const insertTaskService = async (data) => {
    try {
        let task = new Task(data)
        task = await Task.create(task)
        // task =  await task.populate({path: 'createdById', select: '_id name email'})
        return task;
    } catch (error) {
        throw error
    }

}
const getTaskByUserIdService = async (id, dayQuery) => {
    try {
        const user = await User.findById(id);
        const today = dayQuery.startOf("day");

        if(user){
            const tasks = await Task.find({
                createdById: id,
                createdDate: {
                    $gte: today.toDate(),
                    $lte: moment(today).endOf('day').toDate(),
                }
            }) //.populate({path: 'createdById', select: '_id name email'})
            return tasks;
        }
        else{
            throw new Error("User not exist")
        }
    } catch (error) {
        throw error
    }
}
const updateTaskService = async (id, _task) => {
    try {
        task = await Task.findByIdAndUpdate({_id: id}, {
            title: _task?.title,
            description: _task?.description
        }, {new: true});
        if(task){
            return task;
        } 
        else{
            throw Error('Task is not exist')
        } 
    } catch (error) {
        throw error
    }
}
const deleteTaskService = async (id)=>{
    try {
        const count = await Task.deleteOne({_id:id});
        if(count.deletedCount>=1){
            return true;
        }
        else {
            throw new Error("Task is not exist")
        }
    } catch (error) {
        throw error;
    }

}

const getTaskByIdService = async (id)=>{
    try {
        const task = await Task.findById(id)//.populate({path: 'createdById', select: '_id name email'})
        if(task){
            return task;
        }
        else{
            throw new Error("Task is not exist");
        }
    } catch (error) {
        throw error.message;
    }
}

module.exports = {
    insertTaskService,
    getTaskByUserIdService,
    updateTaskService,
    deleteTaskService,
    getTaskByIdService
}