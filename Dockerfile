FROM node:12.22.12-buster-slim
WORKDIR /togo
COPY ./package.json .
COPY . .
RUN mv ./.env-sample ./.env
RUN npm install
EXPOSE 3000:3000
CMD npm run dev
