require('dotenv').config()

const expect  = require('chai').expect;
const user = require('../user');
const task = require('../task');
const redis = require('../async_redis');

var userId;
var taskId;

describe('REDIS', function() {
    it('connection', async function(done){
        this.timeout(5000)
        let redisClient = await redis.asyncCreateClient(process.env.PORT, process.env.REDIS_HOST);
        // console.log(redisClient)
        expect(redisClient).to.be.not.null;
        if(redisClient)
            await redis.asyncQuit(redisClient);
        done();
    })
});


describe('USER', function() {
    it('create user', async function(done){
        this.timeout(5000)

        let mockEvent = {
            headers: {
                Authorization: process.env.AUTH
            },
            body: {
                limit: 2
            }
        }
        mockEvent.body = JSON.stringify(mockEvent.body);

        let result = await user.createUser(mockEvent);
        console.log(result)
        expect(result.statusCode).to.be.eq(200);

        userId = JSON.parse(result.body).id;
        done();
    }),

    it('get user', async function(done){
        this.timeout(5000)

        let mockEvent = {
            headers: {
                Authorization: process.env.AUTH
            },
            body: {
                id: userId
            }
        }
        mockEvent.body = JSON.stringify(mockEvent.body);

        let result = await user.getUser(mockEvent);
        console.log(result)
        expect(result.statusCode).to.be.eq(200);
        
        let data = JSON.parse(result.body).user;
        expect(JSON.parse(result.body).user).to.be.not.null;

        done();
    })
})


describe('Task', function() {
    it('create task', async function(done){
        this.timeout(5000)

        let mockEvent = {
            headers: {
                Authorization: process.env.AUTH
            },
            body: {
                userId: userId,
                description: "test task"
            }
        }
        mockEvent.body = JSON.stringify(mockEvent.body);

        let result = await task.createTask(mockEvent);
        console.log(result)
        expect(result.statusCode).to.be.eq(200);
        expect(JSON.parse(result.body).taskId).to.be.not.null;

        taskId = JSON.parse(result.body).taskId;
        done();
    })

    it('delete user', async function(done){
        this.timeout(5000)

        let mockEvent = {
            headers: {
                Authorization: process.env.AUTH
            },
            body: {
                userId: userId,
                taskId: taskId
            }
        }
        mockEvent.body = JSON.stringify(mockEvent.body);

        let result = await task.deleteTask(mockEvent);
        console.log("Success delete task: " + taskId)
        expect(result.statusCode).to.be.eq(200);
        done();
    })
})
