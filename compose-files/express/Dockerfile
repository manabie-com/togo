FROM node:12

WORKDIR /app
COPY ./compose-files/express/start.sh ./start.sh
COPY . .
RUN npm install
RUN chmod +x ./start.sh

EXPOSE 3200
ENTRYPOINT [ "sh" , "./start.sh" ]