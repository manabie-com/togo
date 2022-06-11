const MDB_TODO = require('../models/MDB_TODO');
const MDB_USER = require('../models/MDB_USER');

module.exports = class TaskClass {
    constructor(data) {
        this.mdb_todo = new MDB_TODO();
        this.mdb_user = new MDB_USER();

        this.data_obj = data;
    }

    async createTask() {
        let res = {};
        try {
            //obj task
            let obj = this.data_obj;
            //find user
            const getUser = await this.mdb_user.findByUser(obj.userName);
            
            //condition: if limit is already 5, you cannot add more task.
            //It will reset evey 8AM via Cron Job    
            if(getUser.limit == 5){
                return res = {
                    message: "You have reached the maximum adding of task per day",
                    data: getUser
                }
            }
            //get the limit and add 1
            let limit = getUser.limit + 1;
            //filter and update for query
            const filter = { username: obj.userName };
            const update = { limit };
            //Will record the task to database
            let data = await this.mdb_todo.add(obj);
            //Will update limit, increment limit
            await this.mdb_user.findOneAndUpdate(filter, update);

            res.status = "success";
            res.data = data;
         
        }catch(error){
            res.status = "error";
            res.message = error.message;
        }

        return res;
    }

    async getTask() {
        let res = {};
        try{
            //Getting all the documents from table todo
            let data = await this.mdb_todo.docs();
            
            res.status = "success";
            res.data = data;

        }catch(error){

            res.status = "error";
            res.message = error.message;
        }

        return res;
    }

    async resetDailyLimit() {
        let res = {};

        try {
            //Will reset the limit to 0 every 8 AM
            await this.mdb_user.findAllAndResetLimit();
            res.status = "success";
        }
        catch (error) {
            res.status = "error";
            res.message = error.message;
        }

        return res;
    }

}