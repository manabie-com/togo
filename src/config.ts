import dotEnv from 'dotenv';

dotEnv.config({ path: '.env' });

export const configuration = (): unknown => ({
  AppSettings: {
    port: process.env.PORT || 8080,
    jwtSecret: process.env.JWT_SECRET
  },
  DbSettings: {
    host: process.env.POSTGRES_HOST,
    port: process.env.POSTGRES_PORT,
    username: process.env.POSTGRES_USERNAME,
    password: process.env.POSTGRES_PASSWORD,
    database: process.env.POSTGRES_DATABASE
  }
});

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export function getConfig<T>(key: string) {
  const config = configuration();
  return (config[key] as T) || ({} as T);
}
