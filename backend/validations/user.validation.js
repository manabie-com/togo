const Joi = require('joi');

const myProfile = {};

const updateLimitDailyTask = {
  body: Joi.object().keys({
    limit_daily_task: Joi.number().required(),
  }),
};

module.exports = {
  myProfile,
  updateLimitDailyTask,
};
