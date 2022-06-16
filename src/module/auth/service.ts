import { Password } from '../../services/password'
import { User } from '../../models/user'
import { BadRequestError } from '../../errors/bad-request-error'
import jwt from 'jsonwebtoken'

interface bodyInterface {
    email: string,
    password: string
}

interface sigInreturn {
    userJwt: string,
    existingUser: object
}

class AuthService {
    
    async signIn(body: bodyInterface) : Promise<sigInreturn> {
        const { email, password } = body
        const existingUser = await User.findOne({email})
    
        if(!existingUser){
            throw new BadRequestError('Invalid creditian')
        }
    
        const isMatch = await Password.comparePassword(existingUser.password, password)
    
        if(!isMatch){
            throw new BadRequestError('Invalid creditian')
        }
    
        const userJwt = jwt.sign({
            id: existingUser.id,
            email: existingUser.email
        }, process.env.JWT_KEY!)

        return { userJwt , existingUser  }
    }

    async signUp(body: bodyInterface) : Promise<sigInreturn> {
        const { email, password } = body
        const existingUser = await User.findOne({ email })
    
        if(existingUser){
            throw new BadRequestError('Email in use')
        }
    
        const user = User.build({
            email,
            password
        })
        await user.save()
        const userJwt = jwt.sign({
            id: user.id,
            email: user.email
        },process.env.JWT_KEY!)

        return { userJwt , existingUser: user  }
    }
}

export const authService = new AuthService()