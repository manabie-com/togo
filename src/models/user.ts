import mongoose from 'mongoose'
import { Password } from '../services/password'

interface UserAttrs {
    email: string,
    password: string
}

export interface UserDoc extends mongoose.Document{
    email: string,
    password: string,
}

interface UserModel extends mongoose.Model<UserDoc>{
    build(attrs: UserAttrs): UserDoc
}


const userSchema = new mongoose.Schema({
    email:{
        type: String,
        required: true
    },
    password:{
        type: String,
    }
}, {
    toJSON: {
        transform(doc, ret){
            ret.id = ret._id
            delete ret._id
            delete ret.__v
            delete ret.password
        }
    }
})


userSchema.statics.build = (attrs: UserAttrs) => {
    return new User(attrs)
}
userSchema.pre('save', async function(done){

    if(this.isModified('password')){
        const hashed = await Password.hashPassword(this.get('password'))
        this.set('password', hashed)
    }
    done
})



const User =mongoose.model<UserDoc, UserModel>('User', userSchema)
export { User }