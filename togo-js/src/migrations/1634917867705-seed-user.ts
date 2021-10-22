import {MigrationInterface, QueryRunner} from "typeorm";

export class seedUser1634917867705 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
      await queryRunner.query(`INSERT INTO "user" ("userId", "password") VALUES ('firstUser', 'example');`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
      // just one value, nothing big deal LOL
      await queryRunner.query(`TRUNCATE "user" RESTART IDENTITY CASCADE;`);
    }

}
