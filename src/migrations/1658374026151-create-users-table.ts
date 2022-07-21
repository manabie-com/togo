import { MigrationInterface, QueryRunner } from "typeorm";

export class createUsersTable1658374026151 implements MigrationInterface {
  name = 'createUsersTable1658374026151'

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`
      CREATE TABLE user_entity(
        id varchar(36) NOT NULL,
        username varchar(255) NOT NULL,
        first_name varchar(255) NOT NULL,
        last_name varchar(255) NOT NULL,
        PRIMARY KEY(id)
      ) ENGINE = InnoDB
    `);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`DROP TABLE user_entity`);
  }

}
