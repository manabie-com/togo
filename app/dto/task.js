const dtoAddTask = (task) => {
    const { status, taskId, userId, name } = task
    return { status, taskId, userId, name }
}

module.exports = {
    dtoAddTask
}