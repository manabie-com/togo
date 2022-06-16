import express, { Request,Response } from 'express'
import { BadRequestError } from '../../errors/bad-request-error'
import { User } from '../../models/user'
import { Password } from '../../services/password'
import jwt from 'jsonwebtoken'
import { authService } from './service'

class AuthController {
    async getCurrentUser(req: Request, res: Response) {
        return res.send({ currentUser: req.currentUser||null })
    }

    async signIn(req: Request, res: Response) {
        const { userJwt, existingUser } = await authService.signIn(req.body)
        req.session = {
            jwt: userJwt
        }
        return res.send(existingUser)
    }

    async signUp(req: Request, res: Response) {
        const { userJwt, existingUser } = await authService.signUp(req.body)
        req.session = {
            jwt: userJwt
        }
        return res.status(201).send(existingUser)
    }

    async logout(req: Request, res: Response) {
        req.session = null
        return res.send({})
    }
}

export const authController = new AuthController()