import { MigrationInterface, QueryRunner } from "typeorm";

export class createSettingsTable1658395075560 implements MigrationInterface {
  name = 'createSettingsTable1658395075560'

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`
    CREATE TABLE setting_entity (
      id varchar(36) NOT NULL,
      todo_per_day int NOT NULL DEFAULT '10',
      user_id varchar(36) NOT NULL,
      UNIQUE INDEX REL_5581db8ead00f4a9cbff338063 (user_id),
      PRIMARY KEY (id)
    ) ENGINE = InnoDB
    `);
    await queryRunner.query(`
    ALTER TABLE
      setting_entity
    ADD
      CONSTRAINT FK_5581db8ead00f4a9cbff3380639 FOREIGN KEY (user_id) REFERENCES user_entity(id) ON DELETE CASCADE ON UPDATE CASCADE
  `);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`ALTER TABLE setting_entity DROP FOREIGN KEY FK_5581db8ead00f4a9cbff3380639`);
    await queryRunner.query(`DROP INDEX REL_5581db8ead00f4a9cbff338063 ON setting_entity`);
    await queryRunner.query(`DROP TABLE setting_entity`);
  }

}
