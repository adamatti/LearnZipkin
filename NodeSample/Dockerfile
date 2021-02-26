FROM node:8.15.0-alpine

USER root
RUN mkdir /app
WORKDIR /app

COPY package.json .
RUN npm install --silent

COPY *.js ./

EXPOSE 3000

CMD npm start