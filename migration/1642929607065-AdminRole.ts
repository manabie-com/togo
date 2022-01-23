import {MigrationInterface, QueryRunner} from "typeorm";
import * as bcrypt from 'bcrypt';
import { Constants } from "src/utils/constants";

export class AdminRole1642929607065 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        const adminInfo = {
            email: 'admin@admin.com',
            password: '1',
            fullName: 'Admin',
            role: Constants.ADMIN_ROLE,
        };
        adminInfo.password = await bcrypt.hash(adminInfo.password, Constants.SALT_OR_ROUNDS);
        await queryRunner.query(`
            insert into \`users\` (email, password, fullName, role) values ('${adminInfo.email}', '${adminInfo.password}' ,'${adminInfo.fullName}', '${adminInfo.role}')
        `);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
    }

}
