`use strict`;

let mysql = require('mysql');
let constant = require('../constant/constant');
let helper = require('../helper/helper');
const Task = require('../models/task');


let con = mysql.createConnection({
  host: "localhost",
  port: "3306",
  user: "root",
  password: "admin",
  database: "demo"
});

//add new task
async function createTask(userID, taskData, date) {
    let taskCount = await getTaskCount(userID, date);
    let newTask;
    if (!userID && !taskData)
    {
      return helper.createResponseObject('Invalid task data', constant.responseFlags.INVALID);
    }
    if (taskCount < constant.taskLimits) {
        newTask = new Task ({
            user_id: userID,
            task_data: taskData,
            date: date,
        });
    } else {
      return helper.createResponseObject('Add task', constant.responseFlags.FULL_TASKS);
    }

    let task = await new Promise((resolve, reject) => {
        newTask.save((err, result) => {
          if (err) return reject(err);
          return resolve(result);
        });
      });
    return task;
}

// get number of task in a specific day
async function getTaskCount(userID, date) {
  return new Promise(async(resolve, reject) => {
    try {
      const values = [userID, date];
      let numOfTask = await con.query('SELECT task_count FROM tb_users WHERE user_id = ? AND date = ?', values);
      return resolve(numOfTask);
    } catch(err) {
      return reject(err);
    }
  });
}

module.exports = {
  createTask
}