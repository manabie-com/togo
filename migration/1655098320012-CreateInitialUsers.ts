import { MigrationInterface, QueryRunner } from "typeorm";

export class CreateInitialUsers1655098320012 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      "INSERT INTO `todo_db`.`users` \
        (`user_id`, `api_key`, `daily_maximum_tasks`) \
      VALUES \
        (10172512, 'aGVsbG8gaXMgaXQgbWUgeW91J3JlIGxvb2tpbmcgZm9yPw==', 4)"
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      "DELETE FROM `todo_db`.`users` WHERE `user_id` in (10172512)"
    );
  }
}
