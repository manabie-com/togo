FROM node:16.13.0-alpine3.14

WORKDIR /usr/src/togo-manabie

ARG NODE_ENV=production
ARG PORT=3000

ENV NODE_ENV $NODE_ENV
ENV PORT $PORT

RUN chown node:node /usr/src/togo-manabie

USER node
COPY --chown=node:node package.json ./

RUN yarn -v

RUN yarn install --prod

COPY --chown=node:node . ./

ENV PORT $PORT
EXPOSE $PORT

CMD [ "ts-node", "--transpile-only", "-r", "tsconfig-paths/register", "src/index.ts"  ]
