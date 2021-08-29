import crypto, { HexBase64Latin1Encoding } from 'crypto';
import pbkdf2_sha512 from 'pbkdf2-sha512';

export class PasswordHelper {

    private static ROUND: number = 25000;
    private static SLAT_SIZE: number = 16;
    private static KEY_LEN_BYTES: number = 64;
    private static ENCODING: HexBase64Latin1Encoding = 'base64';
    private static ALGORITHM: string = 'sha512';

    private static b64trimmed(buf: Buffer): string {
        return buf.toString(PasswordHelper.ENCODING).replace(/=*$/, '').replace('+', '.');
    }

    private static b64decode(str) {
        str = str.replace('.', '+');
        if (str.length % 4) {
            str += '='.repeat(4 - str.length % 4);
        }
        return Buffer.from(str, PasswordHelper.ENCODING);
    }

    private static getHash(password: string, salt: Buffer, rounds: number): string {
        const hash = PasswordHelper.b64trimmed(pbkdf2_sha512(password, salt, rounds, PasswordHelper.KEY_LEN_BYTES));
        return ['', `pbkdf2-${PasswordHelper.ALGORITHM}`, rounds, PasswordHelper.b64trimmed(salt), hash].join('$');
    }

    public static getHmac(secret: string, password: string): string {
        return crypto.createHmac(PasswordHelper.ALGORITHM, secret).update(password).digest(PasswordHelper.ENCODING);
    }

    public static hash(password: string, options?: { saltSize?: number; rounds?: number }) {
        const saltSize: number = options?.saltSize || PasswordHelper.SLAT_SIZE;
        const rounds: number = options?.rounds || PasswordHelper.ROUND;
        const salt: Buffer = crypto.randomBytes(saltSize);
        return PasswordHelper.getHash(password, salt, rounds);
    }

    public static verify(password: string, hashPassword: string) {
        if (!password || !hashPassword) return false;
        const scheme: string = hashPassword.split('$')[1];
        const rounds: number = Number(hashPassword.split('$')[2]) || PasswordHelper.ROUND;
        const salt: string = hashPassword.split('$')[3];
        if (scheme !== `pbkdf2-${PasswordHelper.ALGORITHM}`) return false;
        return hashPassword === PasswordHelper.getHash(password, PasswordHelper.b64decode(salt), rounds);
    }

    static isMediumPassword(password: any): boolean {
        if (password === '') return false;
        return /^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})/.test(password);
    }

    static isStrongPassword(password: any): boolean {
        if (password === '') return false;
        return /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/.test(password);
    }
}
