import { createConnection } from "typeorm";
import config from "../config/ormconfig";

export const createTypeormConn = async () => {
  return await createConnection(config);
};
