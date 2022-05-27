class InvalidUser extends Error {
    errorCode = 'invalidUser'

    constructor(message) {
        super(message);
        this.name = this.constructor.name
    }
}

class LimitTodoPerDay extends Error {
    errorCode = 'limitTodoPerDay'

    constructor(message) {
        super(message);
        this.name = this.constructor.name
    }
}

module.exports = {
    InvalidUser, LimitTodoPerDay
}