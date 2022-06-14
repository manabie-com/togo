import { body } from 'express-validator'

export const signupdaValidator  = [
    body('email')
        .isEmail()
        .withMessage('Email must be valid'),
    body('password')
        .trim()
        .notEmpty()
        .withMessage('Password must be supply')
]
