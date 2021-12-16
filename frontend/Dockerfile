FROM node:16.13.1-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./ui ./

RUN npm install --force
RUN npm install react-scripts@3.4.1 -g

CMD ["npm","start"]