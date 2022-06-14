import { body } from 'express-validator'

export const createTaskValidator  = [
    body().isArray().withMessage("Body is a array"),
    body('*.description')
        .trim()
        .notEmpty()
        .withMessage('Description must be valid'),
    body('*.title')
        .trim()
        .notEmpty()
        .withMessage('Title must be valid'), 
]
