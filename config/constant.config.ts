import { version, name } from "package.json";

/**
 * Configurable of swagger
 */
export const OPTIONS_SWAGGER = {
   apiDocsPath: "/api/v3/docs",
   baseDir: "./",
   exposeApiDocs: true,
   filesPattern: "./**/*.ts",
   info: {
      title: name,
      version,
   },
   security: {
      BasicAuth: {
         scheme: "basic",
         type: "http",
      },
   },
   swaggerUIPath: "/api/docs",
};
