FROM node:14.4.0-alpine as base
WORKDIR /app
COPY ["package.json", "package-lock.json*", "./"]

FROM base as test
RUN npm ci --silent
COPY . .
RUN npm run test:unit
RUN npm run test:itg

FROM base as prod
RUN npm ci --production --silent
RUN npm install -S typescript --silent
COPY . .
RUN npm install -g pm2 --silent
RUN npm run build --silent
EXPOSE $SERVER_PORT
CMD ["pm2-runtime", "process.yaml", "--update-env"]
