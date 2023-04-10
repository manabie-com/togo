const { createTask } = require('./task')
const { MAX_NUMBER_TASK_CREATED, ONE_DAY_TS } = require('../util/constant');
const { User } = require('../model/user');
const { connectDB } = require('../config/db');
const { Task } = require('../model/task');

let today;
let randomEmail
let user;
beforeAll(() => {
    connectDB()
    today = new Date();
    randomEmail = getRandomEmail()
});

const cleanUser = async () => {
    return await User.deleteOne({ _id: user._id })
}

const cleanTask = async () => {
    return await Task.deleteOne({ author: user._id })
}

const getRandomEmail = () => {
    const randomNumber =  Math.floor(Math.random() * 1000).toString();
    return `${randomNumber}@yopmail.com`
}

describe(`Test createTask()`, () => {
    it('Should be success to create task', async () => {
        user = await User.create({
            email: randomEmail,
            password: '123456'
        })

        const newTask = await createTask(user._id, { title: 'hello Manabie', content : 'Hire me pls :)'})
        expect(newTask._id).not.toBe(undefined);
    });
})

afterAll(done => {
    cleanTask().then(resTask => {
        cleanUser().then(resUser => {
            done()
        }).catch(err => {
            done()
        })
    }).catch(err => {
        done()
    })
})