const app = require('../app')
const request = require('supertest')
const {connect, mongoose} = require('../bootstrap/mongoose')
const {init, listDummyUser} = require('../init/user.init')
const User = require('../models/user.model')
const UserSetting = require('../models/userSetting.model')
const ToDo = require('../models/todo.model')

beforeAll(async () => {
    await connect()
    await mongoose.connection.dropDatabase()
    await init()
})
describe('Todo Applications', () => {


    it("should jest work", () => {
        expect(1).toBe(1)
    })

    it("NODE_ENV for testing mockgoose", () => {
        expect(process.env.NODE_ENV).toBe('test')
    })

    it("GET list dummy users", async () => {
        const res = await request(app)
            .get('/users')
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)

        expect(res.body.length).toBe(listDummyUser.length)

        // user props
        expect(res.body).toEqual(expect.arrayContaining([
            expect.objectContaining({
                _id: expect.any(String),
                username: expect.any(String)
            })
        ]))
    })

    it("GET list todo for userId", async () => {
        const getFirstUser = await User.findOne({}, {id: 1}, {})
        const userId = getFirstUser.id
        await ToDo.insertMany([{
            name: 'test1', user: userId
        }, {
            name: 'test2', user: userId
        }])
        const url = `/users/${userId}/todos`
        const response = await request(app)
            .get(url)
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)

        expect(response.status).toBe(200)
        expect(response.body).toEqual(expect.arrayContaining([
            expect.objectContaining({
                name: expect.any(String),
                is_completed: expect.any(Boolean)
            })
        ]))

        const url2 = `/users/123123123/todos`
        const response_2 = await request(app)
            .get(url2)
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)

        expect(response_2.status).toBe(400)
        expect(response_2.body.errorCode).toEqual('invalidUser')
    })

    it("POST a todo for userId", async () => {
        const getFirstUser = await User.findOne({}, {id: 1}, {})
        const userId = getFirstUser.id
        const url = `/users/${userId}/todos`
        const todoName = 'test'
        const response = await request(app)
            .post(url)
            .send({
                name: `${todoName}`
            })
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)

        expect(response.status).toBe(201)
        expect(response.body.name).toEqual(todoName)
    })

    it("POST a todo for userId, invalid userId", async () => {
        const userId = '123712937129837'
        const url = `/users/${userId}/todos`
        const todoName = 'test'
        const response = await request(app)
            .post(url)
            .send({
                name: `${todoName}`
            })
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)

        expect(response.status).toBe(400)
        expect(response.body.errorCode).toEqual('invalidUser')
    })

    it("POST a todo for userId, limit task per day", async () => {
        const getFirstUser = await User.findOne({}, {id: 1}, {})
        const userId = getFirstUser.id
        const userSetting = await UserSetting.findOne({user: userId}, {}, {})
        const limit_per_day = userSetting.limit_per_day

        // init data for user
        const data = []
        for (let i = 0; i < limit_per_day; i++) {
            data.push({name: 'test', user: userId})
        }
        await ToDo.insertMany(data)

        const res = await ToDo.count({user: userId})
        expect(res).toBeGreaterThanOrEqual(50)

        const url = `/users/${userId}/todos`
        const todoName = 'test'
        const response = await request(app)
            .post(url)
            .send({
                name: `${todoName}`
            })
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)

        expect(response.status).toBe(406)
        expect(response.body.errorCode).toEqual('limitTodoPerDay')
    })
})