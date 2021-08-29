import { v1 as uuid } from 'uuid';
import slug from 'slug';
import password from 'generate-password';
import _ from "lodash";

export class StringHelper {
    /**
     * Generate code for sms
     * @param length
     */
    static generateCode(length: number = 4): string {
        if (length < 4) {
            length = 4;
        }

        let code = '';
        for (let i = 0; i < length; i++) {
            code += _.random(0, 9);
        }
        return code;
    }

    /**
     * Generate UUID
     */
    static generateUUID(options?: { prefix?: string; suffix?: string }): string {
        let result: string = uuid();
        if (options?.prefix) result = `${options.prefix}_${result}`;
        if (options?.suffix) result = `${result}_${options.suffix}`;
        return result;
    }

    static generatePassword(): string {
        return password.generate({
            length: 16,
            numbers: true,
            symbols: true,
            uppercase: true,
            lowercase: true,
            excludeSimilarCharacters: true,
            strict: true
        });
    }

    /**
     * Create a slug from list of string
     * @param values
     */
    static slugify(...values: string[]): string {
        let result: string = '';
        values.map((value) => {
            if (!!value) {
                result += '-' + slug(value, {lower: true});
            }
        });

        if (result) {
            return result.substr(1);
        }

        return '_';
    }

    static numberThousand(num: number, separator: 'dot' | 'commas' = 'dot'): string {
        const numStr: string = num ? num.toString() : '0';
        const separatorChar = separator === 'dot' ? '.' : ',';
        return numStr.replace(/\B(?=(\d{3})+(?!\d))/g, separatorChar);
    }

    static pattern(template: string, ...params: string[]): string {
        let pattern: string = template;
        for (let i = 0; i < params.length; i++) {
            pattern = pattern.replace(`{${i}}`, params[i]);
        }
        return pattern;
    }
}
