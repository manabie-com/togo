let taskUsers = [];

module.exports.find = (user_id, date) => {
    return taskUsers.filter((task) => task.user_id === user_id && task.date === date)
};

module.exports.save = ({user_id, task_name, date}) => {
    user = {
        user_id, task_name, date
    };
    taskUsers.push(user);
    return user;
};