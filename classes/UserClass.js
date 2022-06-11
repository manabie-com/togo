const MDB_USER = require('../models/MDB_USER');
const bcrypt = require("bcryptjs") //to encrypt password
const jwt = require("jsonwebtoken") //generate token


module.exports = class TaskClass {
    constructor(data) {
        this.mdb_user = new MDB_USER();

        this.data_obj = data;
    }

    async validate() {
        let res = {};
        let obj = this.data_obj;
        // console.log(obj.username);
        // let check_user = await this.mdb_user.findByUser(obj.username);
        // console.log(check_user.username);
        // console.log(check_user.username == obj.username);
        if (obj.fullname.trim() == '' || obj.confirmpassword.trim() == '' || obj.password.trim() == '' || obj.username.trim() == '') {
            res.status = "error";
            res.message = "You need to fill up all fields in order to proceed.";
        }
        else if (obj.password.length < 8 || obj.confirmpassword.length < 8) {
            res.status = "error";
            res.message = "Password must be more than 8 characters.";
        }
        else if (obj.password.length > 16 || obj.confirmpassword.length > 16) {
            res.status = "error";
            res.message = "Password must not be more than 16 characters.";
        }
        else if (obj.confirmpassword !== obj.password) {
            res.status = "error";
            res.message = "The password you entered didn't match.";
        }
        // else if (check_user.username == obj.username) {
        //     res.status = "error";
        //     res.message = "The username you entered already exists.";
        // }
        else {
            res.status = "success";
            res.data = obj;
        }

        return res;
    }

    async create() {
        let res = {};
        let obj = this.data_obj;
        let hashed_password = '';
        hashed_password = await bcrypt.hash(obj.password, 10);
        
        try {

            let add_form =
            {
                fullname: obj.fullname,
                password: hashed_password,
                username: obj.username
            }
            let data = await this.mdb_user.add(add_form);
            
            res.data = data;


            res.status = "successfuly registered";
        }
        catch (error) {
            console.log(error);
            res.status = "error";
            res.message = error.message;
        }

        return res;
    }

    async authenticate() {
        let res = {};
        let obj = this.data_obj;

        if (obj.password.trim() == '' || obj.username.trim() == '') {
            res.status = "error";
            res.message = "You need to fill up all fields in order to proceed.";
        }
        else {
            console.log(obj.username);
            let check_account = await this.mdb_user.findByUser(obj.username);
            console.log(check_account);
            if (check_account) {
                const result = await bcrypt.compare(obj.password, check_account.password);
                if (!result) {
                    res.status = "error";
                    res.message = "Invalid Credentials";
                }
                else {
                    check_account.token = jwt.sign(obj, 'akaru-todo', { expiresIn: '1h' })
                    res.status = "success";
                    res.data = check_account;
                }
            }
            else {
                res.status = "error";
                res.message = "Invalid Credentials";
            }
        }

        return res;
    }

}