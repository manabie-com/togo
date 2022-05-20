import * as bcrypt from 'bcrypt';

export class CryptoUtil {
    static async hash(value: string, saltOrRounds: number = 10): Promise<string> {
		  return await bcrypt.hash(value, saltOrRounds);
    }

    static async compareHashWithPlainText(hashValue: string, plainText: string): Promise<boolean> {
		  return await bcrypt.compare(plainText, hashValue);
    }
}