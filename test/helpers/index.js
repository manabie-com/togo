const bcrypt = require('bcryptjs')
const request = require('supertest')
const config = require('../../src/config')
const dbConfig = require('../../src/config/database')
const Sequelize = require('sequelize')
const sequelize = new Sequelize(dbConfig)
const app = require('../../src/app')
const SALT = 8
const API_SIGNUP = `${config.API_BASE}/auth/signup`
const API_SIGNIN = `${config.API_BASE}/auth/signin`

const USER_TEST = {
    name: 'User test',
    email: `user@test.com`,
    password: '12345678',
    role: 'Admin'
}

const USER_SIGNIN = {
    email: `user@test.com`,
    password: '12345678'
}

async function signup() {
    //await startDatabase()
    await request(app).post(API_SIGNUP).send(USER_TEST)
}


async function signin() {
    const query = `SELECT * FROM "users" where email like 'user@test.com'`
    const user = await sequelize.query(query)
    console.log('user', user[0])
    if (user[0].length <= 0) {
        await signup()
        const response = await request(app).post(API_SIGNIN).send(USER_TEST)
        return response.headers['set-cookie'][0]
    } else {
        const response = await request(app).post(API_SIGNIN).send(USER_SIGNIN)
        return response.headers['set-cookie'][0]
    }

}

const helper = {
    USER_TEST,
    signin
}

module.exports = helper