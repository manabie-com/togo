FROM node:14.17.6
ARG PORT
ARG NODE_ENV
ARG JWT_SECRET
ENV PORT=${PORT}
ENV NODE_ENV=${NODE_ENV}
ENV JWT_SECRET=${JWT_SECRET} 
WORKDIR /app
COPY ["package.json", "package-lock.json*", "./"]
RUN npm install
COPY . .
CMD ["npm", "start"]