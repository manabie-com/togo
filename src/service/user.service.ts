import { injectable } from "inversify";
import ms from 'ms';
import jwt from "jsonwebtoken";

@injectable()
export class UserService {

    public async createToken(userId: string): Promise<string> {
        const secretKey: string = String(process.env.JWT_AUTH_SECRET_KEY);
        const expiresInSecond: number = ms(String(process.env.JWT_AUTH_EXPIRATION)) / 1000;
        const payload: any = {user_id: userId};
        return jwt.sign(payload, secretKey, {expiresIn: expiresInSecond});
    }
}
