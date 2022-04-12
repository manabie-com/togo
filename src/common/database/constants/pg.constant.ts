export class PgConstants {
  // POSTGRES
  public static readonly POSTGRES = 'postgres';

  // POSTGRES MASTER
  public static readonly POSTGRES_HOST = 'POSTGRES_HOST';
  public static readonly POSTGRES_PORT = 'POSTGRES_PORT';
  public static readonly POSTGRES_USER = 'POSTGRES_USER';
  public static readonly POSTGRES_PASSWORD = 'POSTGRES_PASSWORD';
  public static readonly POSTGRES_DATABASE = 'POSTGRES_DATABASE';

  // TYPEORM
  public static readonly TYPE_ORM_ENTITIES = 'dist/**/*.entity{.ts,.js}';
  public static readonly TYPE_ORM_MIGRATION_TABLE_NAME = 'migration';
  public static readonly TYPE_ORM_MIGRATIONS = 'dist/migrations/*.js';
  public static readonly TYPE_ORM_CLI_MIGRATIONS_DIR = 'src/migrations';
}
