import UserModel from '../Models/Postgres/Users'
import database from "../Connectors"
import {Sequelize} from 'sequelize'
import md5 from 'md5'
class UserRepositories {

    constructor() {
        UserModel.init(database, Sequelize)
    }

    getUser = async (params) => {
        let {username, password, customFields, userId} = {...params}
        let defaultFields = [
            "user_id",
            "username",
            "max_todo"
        ]
        if(customFields) defaultFields = customFields
        let condition = {}
        if(userId){
            condition.user_id = userId
        }
        if(username && password){
            password = md5(md5(password) + username)
            condition.username = username
            condition.password = password
        }
        let user = []

        try{
            user = await UserModel.findOne({
                attributes: defaultFields,
                where: condition
            })

        }catch (e) {
            return null
        }
        return user
    }

    updateUser = async (params, userId = 0) => {
        try{
            let user
            if (userId === 0) {
                user = await UserModel.create(params)
            } else {
                user = await UserModel.update(params, {
                    where: {
                        user_id: userId
                    }
                })
            }
            return user
        }catch (e) {
            return false
        }
    }

    deleteUser = async (params) => {
        try{
            await UserModel.destroy({
                where: { user_id: params.userId }
            })
        }catch (e) {
            return false
        }
        return true
    }

}

export default new UserRepositories()
