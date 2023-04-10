const { checkUserQuota, updateUserQuota } = require('./user')
const { MAX_NUMBER_TASK_CREATED, ONE_DAY_TS } = require('../util/constant');
const { User } = require('../model/user');
const { connectDB } = require('../config/db')

let today;
let randomEmail
beforeAll(() => {
    connectDB()
    today = new Date();
    randomEmail = getRandomEmail()
});

const cleanUser = async () => {
    return await User.deleteOne({ email: randomEmail })
}

const getRandomEmail = () => {
    const randomNumber =  Math.floor(Math.random() * 1000).toString();
    return `${randomNumber}@yopmail.com`
}

describe(`Test checkUserQuota()`, () => {
    const sleep = ms => new Promise(r => setTimeout(r, ms));

    it('Should be success when user create first task', async () => {
        const quota = {
            max_post_by_day: MAX_NUMBER_TASK_CREATED,
            last_task_created_at: undefined,
            remaining_post: MAX_NUMBER_TASK_CREATED
        }
        let taskCreatedAt = new Date(today.getFullYear(), today.getMonth(), today.getDate())
        const { newRemainingPost, createdAt } = checkUserQuota(quota);
        expect(newRemainingPost).toBe(MAX_NUMBER_TASK_CREATED - 1);
        expect(Math.floor(taskCreatedAt.getTime() / 1000)).toBe(Math.floor((createdAt.getTime() / 1000)));
    });

    it('Should deduct 1 task if quota is enough within 1 day', async () => {
        // const twoDayBeforeToday = new Date(Math.floor(today.getTime() / 1000) - 2 * ONE_DAY_TS)
        let taskCreatedAt = new Date(today.getFullYear(), today.getMonth(), today.getDate())
        const quota = {
            max_post_by_day: MAX_NUMBER_TASK_CREATED,
            last_task_created_at: taskCreatedAt,
            remaining_post: MAX_NUMBER_TASK_CREATED - 1
        }
        const { newRemainingPost, createdAt } = checkUserQuota(quota);
        expect(newRemainingPost).toBe(quota.remaining_post - 1);
        expect(createdAt).toBe(undefined);
    });


    it('Should refill quota if newly created task is different from the last task in db ', async () => {
        const twoDayBeforeToday = new Date(Math.floor(today.getTime() / 1000) - 2 * ONE_DAY_TS)
        let taskCreatedAt = new Date(today.getFullYear(), today.getMonth(), today.getDate())
        const quota = {
            max_post_by_day: MAX_NUMBER_TASK_CREATED,
            last_task_created_at: twoDayBeforeToday,
            remaining_post: 0
        }
        const { newRemainingPost, createdAt } = checkUserQuota(quota);
        expect(newRemainingPost).toBe(MAX_NUMBER_TASK_CREATED - 1);
        expect(Math.floor(taskCreatedAt.getTime() / 1000)).toBe(Math.floor((createdAt.getTime() / 1000)));
    });
})

describe(`Test updateUserQuota()`, () => {
    it('Should be success when user create first task', async () => {
        user = await User.create({
            email: randomEmail,
            password: '123456'
        })

        const newRemainingPost = MAX_NUMBER_TASK_CREATED - 2
        const taskCreatedAt = new Date(today.getFullYear(), today.getMonth(), today.getDate())
        await updateUserQuota(user, newRemainingPost, taskCreatedAt)
        expect(user.quota.remaining_post).toBe(newRemainingPost);
        expect(Math.floor(taskCreatedAt.getTime() / 1000)).toBe(Math.floor((user.quota.last_task_created_at.getTime() / 1000)));
    });
})

afterAll(done => {
    cleanUser().then(resUser => {
        done()
    }).catch(err => {
        done()
    })
})