FROM node:16-alpine

WORKDIR /app

COPY ["package.json", "package-lock.json", "./"]

RUN npm install --silent

COPY . .

EXPOSE 3002

CMD ["npm", "start"]