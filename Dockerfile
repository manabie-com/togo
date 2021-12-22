FROM node:14.0.0
ARG PORT
ARG MONGODB_URL
ARG JWT_SECRET
ENV PORT=${PORT}
ENV MONGODB_URL=${MONGODB_URL}}
ENV JWT_SECRET=${JWT_SECRET} 
WORKDIR /app
COPY ["package.json", "package-lock.json*", "./"]
RUN npm install
COPY . .
CMD ["npm", "start"]