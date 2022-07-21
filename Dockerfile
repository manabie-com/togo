FROM node:16.14.0-stretch AS base

WORKDIR /opt/app

FROM base AS build
COPY package.json package-lock.json* ./
RUN npm install && npm cache clean --force
COPY . ./
RUN npm run build

FROM node:16.14-alpine as release
WORKDIR /opt/app
COPY --from=build /opt/app/package.json /opt/app/package-lock.json* ./
RUN npm install --only-production
COPY --from=build /opt/app/dist ./dist

ENV NODE_OPTIONS=--max_old_space_size=4096
CMD ["node", "dist/main.js"]
