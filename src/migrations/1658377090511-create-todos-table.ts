import { MigrationInterface, QueryRunner } from "typeorm";

export class createTodosTable1658377090511 implements MigrationInterface {
  name = 'createTodosTable1658377090511'

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`
      CREATE TABLE todo_entity (
        id varchar(36) NOT NULL,
        title varchar(255) NOT NULL,
        date date NOT NULL,
        userId varchar(36) NULL,
        PRIMARY KEY (id)
      ) ENGINE = InnoDB
    `);
    await queryRunner.query(`
      ALTER TABLE
        todo_entity
      ADD
        CONSTRAINT FK_f3037daa47e75647225318cc58e FOREIGN KEY (userId) REFERENCES user_entity(id) ON DELETE CASCADE ON UPDATE CASCADE
    `);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`ALTER TABLE todo_entity DROP FOREIGN KEY FK_f3037daa47e75647225318cc58e`);
    await queryRunner.query(`DROP TABLE todo_entity`);
  }

}
