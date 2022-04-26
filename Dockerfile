FROM node:alpine

RUN mkdir -p manabie-todoapp && chown -R node:node manabie-todoapp

WORKDIR manabie-todoapp

COPY . .

USER node

RUN npm install

COPY --chown=node:node . .

EXPOSE 6363

CMD ["npm","start"]
