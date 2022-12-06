let userList = require('./data/users.json');
let taskList = require('./data/tasks.json');
const fs = require('fs');

module.exports = {
    create: (req, res) => {
        const { username = '', date } = req.body;
        const index = userList.findIndex((item) => item.username === username);
        const user = userList[index];
        if (user) {
            const taskListByUser = taskList.filter((item) => item.username === username && item.date === date);
            if (taskListByUser.length === user.limitTaskNumber) {
                res.status(400).json({
                    message: "Your task is fully for today!"
                });
                return;
            }
            taskList.push(req.body);
            const filePath = `${__dirname}\\data\\tasks.json`;
            fs.writeFileSync(filePath, JSON.stringify(taskList));
            res.status(201).json({
                message: "Your task have been created!"
            });
        }
        
    }
}