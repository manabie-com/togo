import TaskModel from '../Models/Postgres/Tasks'
import database from "../Connectors"
import {Sequelize, Op} from 'sequelize'
class TaskRepositories {

    constructor() {
        TaskModel.init(database, Sequelize)
    }

    getTasks = async (params) => {
        let {content, dateFrom, dateTo, taskId, userId, onlyCount, customFields} = {...params}
        let defaultFields = [
            "task_id"
        ]
        if(customFields) {
            defaultFields = customFields
        }

        let condition = {}

        if(userId){
            if(!Array.isArray(userId)) userId = userId.split(",")
            condition.user_id = {
                [Op.in]: userId
            }
        }

        if(taskId){
            if(!Array.isArray(taskId)) taskId = taskId.split(",")
            condition.task_id = {
                [Op.in]: taskId
            }
        }
        if(dateFrom && dateTo){
            condition.created_date = {
                [Op.between]: [dateFrom, dateTo]
            }
        }else{
            if(dateFrom){
                condition.created_date = {
                    [Op.gte]: dateFrom
                }
            }
            if(dateTo){
                condition.created_date = {
                    [Op.lte]: dateTo
                }
            }
        }


        if(content){
            condition.content = {
                [Op.like]: `%${content}%`
            }
        }

        let tasks = []

        try{
            if(onlyCount){
                return await TaskModel.count({
                    where: condition
                })
            }
            tasks = await TaskModel.findAll({
                attributes: defaultFields,
                where: condition
            })
        }catch (e) {
            return null
        }
        return tasks
    }

    updateTask = async (params, taskId = 0) => {
        try{
            let task
            if (taskId === 0) {
                task = await TaskModel.create(params)
            } else {
                task = await TaskModel.update(params, {
                    where: {
                        task_id: taskId
                    }
                })
            }
            return task
        }catch (e) {
            return false
        }
    }

    deleteTask = async (params) => {
        try{
            await TaskModel.destroy({
                where: { task_id: params.taskId }
            })
        }catch (e) {
            return false
        }
        return true
    }

}

export default new TaskRepositories()
