import * as bcrypt from 'bcryptjs';

export class EncryptionUtil {
  public static PasswordEncryptionDefaultSaltLength = 11;

  public static async encryptPassword(
    password: string,
    saltLength: number = this.PasswordEncryptionDefaultSaltLength,
  ): Promise<string> {
    return await bcrypt.hash(password, saltLength);
  }

  public static comparePassword(password: string, hash: string) {
    return bcrypt.compareSync(password, hash);
  }
}
