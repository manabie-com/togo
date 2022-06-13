import { getConnection } from "typeorm";

export const getTypeormConn = async () => {
  return await getConnection();
};
