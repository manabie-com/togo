const { IN_SUFFICIENT_QUOTA, ONE_DAY_TS } = require('../util/constant');
const { User } = require('../model/user');
const { storeLog } = require('../util/log')

/**
 * Check if current time is different date from user's `last_task_created_at` on db
 * @param {Number} nowTs - current timestamp in seconds
 * @param {*} quota - user quota
 * @returns true if same day otherwise false
 */
const isSameDay = (nowTs, quota) => {
    const lastTaskTime = Math.floor(quota.last_task_created_at.getTime()/ 1000);
    // console.log(nowTs, lastTaskTime);
    return nowTs - lastTaskTime < ONE_DAY_TS
}

const userService = {
    createUser: async (payload) => {
        return await User.create(payload)
    },
    /**
     * Check if user have enough quota to create task
     * @param {Object} quota - user quota
     * @param {Number} quota.max_post_by_day - maximum number of task user can create per day
     * @param {Number} quota.remaining_post - number of remaining task available within a day
     * @param {Date} quota.last_task_created_at - specific date time at 00:00:00 (hh:mm:ss)
     * @returns return newRemainingPost is number of task remaining (after this task is created), createdAt(only return if current last_task_created_at is outdate or undefined) is date time of current time of creating task
     */
    checkUserQuota: (quota) => {
        let newRemainingPost;
        const today = new Date();
        let createdAt = new Date(today.getFullYear(), today.getMonth(), today.getDate())
        if (!quota.last_task_created_at) {
            newRemainingPost = quota.max_post_by_day - 1;
            return { newRemainingPost, createdAt }
        }
        const nowTs = Math.floor(Date.now() / 1000)
        const isSameDayPost = isSameDay(nowTs, quota);
        if (isSameDayPost && quota.remaining_post === 0) {
            return { newRemainingPost: IN_SUFFICIENT_QUOTA };
        }
        if (isSameDayPost) {
            // same day so no need to update create time
            newRemainingPost = quota.remaining_post - 1
            return { newRemainingPost };
        } else {
            // update new task creation time
            newRemainingPost = quota.max_post_by_day - 1
            return { newRemainingPost, createdAt };
        }
    },
    /**
     * Update new user quota after task was successfully created
     * @param {Object} user - Mongoose User object
     * @param {Number} newRemainingPost - new remaining number of task available within a day
     * @param {Date} createdAt - specific date time at 00:00:00 (hh:mm:ss)
     */
    updateUserQuota: async (user, newRemainingPost, createdAt) => {
        try {
            user.quota.remaining_post = newRemainingPost;
            if (createdAt) {
                user.quota.last_task_created_at = createdAt;
            }
            await user.save();
        } catch (error) {
            // store log and notify for mannual fix
            storeLog({
                type: 'ERROR',
                source: 'updateUserRemainingPost',
                description: `Could not deduct user remaining task creation ${user.email}`,
                error
            })
        }
    }
}

module.exports = userService;
