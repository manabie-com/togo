const Joi = require("joi");

const { password } = require("./customize.validation");

const registerSchema = {
  body: Joi.object().keys({
    name: Joi.string().required(),
    email: Joi.string().required().email(),
    password: Joi.string().required().custom(password),
    maxTask: Joi.number().required().min(1),
  }),
};

const loginSchema = {
  body: Joi.object().keys({
    email: Joi.string().required(),
    password: Joi.string().required(),
  }),
};

module.exports = {
  registerSchema,
  loginSchema,
};
