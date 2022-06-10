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
            console.log(obj.userName);
            //find user
            const getUser = await this.mdb_user.findByUser(obj.userName);
            //get the limit and add 1
            console.log(getUser.limit);
            let limit = getUser.limit + 1;
            console.log(limit);
            const filter = { userName: obj.userName };
            const update = { limit };

            let data = await this.mdb_todo.add(obj);
            await this.mdb_user.findOneAndUpdate(filter, update);

            res.status = "success";
            res.data = data;
         
        }catch(error){
            console.log(error);
            res.status = "error";
            res.message = error.message;
        }

        return res;
    }

    async getTask() {
        let res = {};
        try{
            // let obj = this.data_obj;
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