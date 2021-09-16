interface IAppSettings {
  readonly port: string;
  readonly jwtSecret: string;
}

interface IDBSettings {
  readonly host: string;
  readonly port: number;
  readonly username: string;
  readonly password: string;
  readonly database: string;
}
