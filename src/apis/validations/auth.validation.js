const Joi = require("joi");

const { password } = require("./customize.validation");

const registerSchema = {
  body: Joi.object().keys({
    name: Joi.string().required().min(5).max(50),
    email: Joi.string().required().min(5).max(255).email(),
    password: Joi.string().required().min(5).max(1024).custom(password),
    maxTask: Joi.number().required().min(1),
  }),
};

const loginSchema = {
  body: Joi.object().keys({
    email: Joi.string().required().min(5).max(255).email(),
    password: Joi.string().required().min(5).max(1024).custom(password),
  }),
};

module.exports = {
  registerSchema,
  loginSchema,
};
