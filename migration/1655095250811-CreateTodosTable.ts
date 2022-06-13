import { MigrationInterface, QueryRunner } from "typeorm";

export class CreateTodosTable1655095250811 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      "CREATE TABLE IF NOT EXISTS `todos` (\
        `id` int (11) NOT NULL AUTO_INCREMENT, \
        `task` text NOT NULL, \
        `user_id` int (11), \
        `creation_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, \
        PRIMARY KEY (`id`) \
      )"
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query("DROP TABLE IF EXISTS `todos`;");
  }
}
