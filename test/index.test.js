const request = require('supertest');
const app = require('../app')
const { MAX_NUMBER_TASK_CREATED, ERROR } = require('../util/constant')
const { connectDB, getDBConn } = require('../config/db')
const { User } = require('../model/user')
const { Task } = require('../model/task')
/* Connecting to the redis at first*/

const getRandomEmail = () => {
    const randomNumber =  Math.floor(Math.random() * 1000).toString();
    return `${randomNumber}@yopmail.com`
}

let randomEmail;
let secondRandomEmail
const userIds = []
let conn;
beforeAll(() => {
    conn = connectDB()
    randomEmail = getRandomEmail()
    secondRandomEmail = getRandomEmail()
});

const cleanUsers = async () => {
    return await User.deleteMany({ email: [randomEmail,  secondRandomEmail]})
}

const cleanTasks = async () => {
   return await Task.deleteMany({ author: userIds })
}

const password = '123456'
describe(`Test User API`, () => {
    it(`Create user successfully`, async () => {
        const res = await request(app)
            .post("/user/register")
            .send({ email: randomEmail , password });
        expect(res.statusCode).toBe(201);
        userIds.push(res.body._id)
    });

    it(`Should loggin sucessfully`, async () => {
        const res = await request(app)
        .post("/user/login")
        .send({ email: randomEmail , password })
        expect(res.statusCode).toBe(200);
    });
});


/* Testing the API endpoints. */
describe(`Test Task API`, () => {
    it(`Should create task successfully`, async () => {
        const loginResult = await request(app)
        .post("/user/login")
        .send({ email: randomEmail , password })

        const res = await request(app)
            .post("/task")
            .set('Authorization', `Bearer ${loginResult.body.access_token}`)
            .send({
                title: 'title',
                content: 'content 123'
            })
        expect(res.statusCode).toBe(201);
    });

    it(`Should hit rate limit with more than ${MAX_NUMBER_TASK_CREATED} tasks were created`, async () => {

        // start preparation data for test
        // user registration
        const newUser = await request(app)
        .post("/user/register")
        .send({ email: secondRandomEmail , password })
        userIds.push(newUser.body._id)

        // new user login
        const loginResult = await request(app)
        .post("/user/login")
        .send({ email: secondRandomEmail , password })


        const createRandomTask = async () => {
            const createTaskResult = await request(app)
            .post("/task")
            .set('Authorization', `Bearer ${loginResult.body.access_token}`)
            .send({
                title: 'title1',
                content: 'content 1234'
            })
            return createTaskResult;
        }
        // use all quota in 1 day of a user
        for (let i = 0; i < MAX_NUMBER_TASK_CREATED; i++) {
            const successCreateTaskRes = await createRandomTask()
            expect(successCreateTaskRes.statusCode).toBe(201);
        }
        // end preparation data for test
        // start test
        const failedCreateTask = await createRandomTask()
        expect(failedCreateTask.statusCode).toBe(ERROR.IN_SUFFICIENT_QUOTA.status_code);
        expect(failedCreateTask.body.error_msg).toBe(ERROR.IN_SUFFICIENT_QUOTA.error_msg);
    });
});

/* Close redis connection */
afterAll(done => {
    cleanTasks().then(resTask => {
        cleanUsers().then(resUser => {
            done()
        }).catch(err => {
            done()
        })
    }).catch(err => {
        done()
    })
})