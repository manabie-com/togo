const User = require('../models/user.model')
const UserSetting = require('../models/userSetting.model')
const bcrypt = require("bcrypt");
const defaultPasswd = 'password'

const dummyUsers = [
    {
        username: 'user_1',
        limit_per_day: 50
    },
    {
        username: 'user_2',
        limit_per_day: 50
    },
    {
        username: 'user_3',
        limit_per_day: 100
    },
    {
        username: 'user_4',
        limit_per_day: 5000
    },
    {
        username: 'user_5',
        limit_per_day: 5000
    },
    {
        username: 'user_5',
        limit_per_day: 5000
    }
]

const createDummyUser = async () => {
    const preparedUserDoc = dummyUsers.map(ele => new User({
        username: ele.username, password: bcrypt.hashSync(defaultPasswd, 5)
    }))
    return await User.insertMany(preparedUserDoc)
}

const createDummyUserSetting = async (users) => {
    const preparedDocs = users.map(ele => {
            const findDummySetting = dummyUsers.find(dummyEle => dummyEle.username === ele.username)
            const limit_per_day = findDummySetting ? findDummySetting.limit_per_day : 10
            return new UserSetting({
                user: ele,
                limit_per_day: limit_per_day
            })
        }
    )
    return await UserSetting.insertMany(preparedDocs)
}

const init = async () => {
    try {
        const numUser = await User.count({})
        if (!numUser) {
            console.log("init dummy users and user settings")
            const users = await createDummyUser()
            const res = await createDummyUserSetting(users)
            return `Created ${res.length} dummy users`
        } else {
            return `There are ${numUser} dummy users`
        }
    } catch (err) {
        throw err
    }

}

module.exports = {
    init: async () => {
        console.log(await init())
    },
    listDummyUser: dummyUsers
}


