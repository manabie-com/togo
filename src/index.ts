import mongoose from 'mongoose'
import { app }  from './app'
import * as dotenv from 'dotenv';

dotenv.config({ path: __dirname+'/LCL.env' });

const start = async () =>{
    try{
        if(!process.env.JWT_KEY){
            throw new Error('JWT_KEY must be define')
        }
        if(!process.env.MONGO_URI){
            throw new Error('MONGO_URI must be define')
        }
        await mongoose.connect(process.env.MONGO_URI,{
            useFindAndModify: true,
            useNewUrlParser: true,
            useUnifiedTopology: true
        })
        console.log('Connect MongoDB!!')
    }catch(err){
        console.error(err)
    }
    app.listen(3000,()=>{
        console.log(`Server is running at 3000!!!`)
    })
}

start()



