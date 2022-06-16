import { body } from 'express-validator'

export const signinValidator  = [
    body('email')
        .isEmail()
        .withMessage('Email must be valid'),
    body('password')
        .trim()
        .notEmpty()
        .withMessage('Password must be supply')
]
