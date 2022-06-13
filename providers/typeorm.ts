import { Connection, ConnectionOptions, createConnections } from "typeorm";
import { SnakeNamingStrategy } from "typeorm-naming-strategies";
import dbconfig from "../config/ormconfig";

const defaultOptions: any = {
  ...dbconfig,
  namingStrategy: new SnakeNamingStrategy(),
};

class Database {
  private _connection: Connection = null;

  async initialize(options?: ConnectionOptions): Promise<Connection> {
    try {
      options = options || defaultOptions;
      const [defaultConnection] = await createConnections([
        {
          name: options.name || "someConnectionName",
          ...options,
        },
      ]);

      this._connection = defaultConnection;

      return this._connection;
    } catch (error) {
      console.log(error);
    }
  }

  getConnection(): Connection {
    return this._connection;
  }

  async stop() {
    await this._connection.close();
  }
}

export default new Database();
