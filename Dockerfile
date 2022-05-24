FROM node:16-alpine

WORKDIR /togo/


COPY package.json ./package.json
COPY package-lock.json ./package-lock.json


RUN apk add curl && \
    apk add --update --no-cache curl py-pip make alpine-sdk && \
    npm install && \
    apk del python3 make alpine-sdk && \
    rm /var/cache/apk/* && \
    rm -rf /root/.npm /root/.node-gyp && \
    rm -rf /usr/lib/node_modules && \
    rm -rf /tmp/*

COPY ./src ./src

ENV NODE_ENV production
ENV PORT 9200

EXPOSE 9200
CMD [ "npm", "start" ]