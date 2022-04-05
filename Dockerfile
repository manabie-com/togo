FROM node:12-alpine as fromBuilder

WORKDIR /app
COPY . /app
RUN npm install
RUN npm run build

FROM node:12-alpine

WORKDIR /home/app

COPY --from=fromBuilder /app/dist /home/app/dist
COPY --from=fromBuilder /app/node_modules /home/app/node_modules

CMD [ "node", "dist/main.js" ]
