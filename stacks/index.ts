import * as sst from "@serverless-stack/resources";
import * as cdk from "aws-cdk-lib";

import { CoreStack } from "./CoreStack";

export default function (app: sst.App) {
  // Set default runtime for all functions
  app.setDefaultFunctionProps({
    runtime: "go1.x",
    srcPath: "backend",
    logRetention:
      app.stage === "prod"
        ? cdk.aws_logs.RetentionDays.THREE_MONTHS
        : cdk.aws_logs.RetentionDays.ONE_DAY,
    tracing: !app.local ? "active" : "disabled",
    environment: {
      APP_STAGE: app.stage,
    },
  });

  // Remove all resources when non-prod stages are removed
  if (app.stage !== "prod") {
    app.setDefaultRemovalPolicy("destroy");
  }

  app.stack(CoreStack);
}

export function debugApp(app: sst.App) {
  // Make sure to create the DebugStack when using the debugApp callback
  new sst.DebugStack(app, "DebugStack");
}
