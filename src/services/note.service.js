const noteModel = require('../models/note.model')
const moment = require("moment");

module.exports = {
  checkLimitTask: async function (id) {
  try {
    const today = moment().startOf("day");
    let countTask = await noteModel.count({
      user: id,
      deleted: false,
      createdAt: {
        $gte: today.toDate(),
        $lte: moment(today).endOf('day').toDate(),
      },
    });
    // console.log(countTask, moment("2021-12-03").endOf("day").toDate());
    if (countTask >= 5) return true;
    return false;
  } catch (err) {
    throw new Error(err.message);
  }
},
    
  addTask2DB: async function (content, id) {
    try {
      const newTask = await noteModel.create({
        content: content,
        user: id,
      });
      return newTask;
    } catch (err) {
      throw new Error(err.message);
    }
  },

  makeTaskCompleted: async function (id, userId) {
    try {
      const tickTask = await noteModel.findOneAndUpdate({
        _id: id,
        user: userId,
        deleted: false,
        ticked: true,
      });
      return tickTask;
    } catch (err) {
      throw new Error("Opps, something went wrong");
    }
  },

  getTaskList: async function (userId, day) {
    try {
      const queryDay = moment(day).startOf("day");
      let taskList = await noteModel
        .find(
          {
            user: userId,
            deleted: false,
            createdAt: {
              $gte: queryDay.toDate(),
              $lte: moment(queryDay).endOf("day").toDate(),
            },
          },
         
        );
      return taskList;
    } catch (err) {
      throw new Error(err.message);
    }
  },

  updateTask: async function (id, userId, content, ticked ) {
    try {
      const taskUpdated = noteModel.findOneAndUpdate({
        _id: id,
        deleted: false,
        user: userId
      },{
        content: content,
        ticked: ticked
      },{ new: true})
      return taskUpdated;
    } catch (err) {
      throw new Error(err.message)
    }
  },

  deleteTask: async function (id, userId) {
    try {
      const taskDeleted = noteModel.findOneAndUpdate({
        _id: id,
        deleted: false,
        user: userId
      },{
        deleted: true
      },{ new: true})
      return taskDeleted;
    } catch (err) {
      throw new Error(err.message)
    }
  }
};
