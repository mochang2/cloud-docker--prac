FROM node:alpine

# WORKDIR must precede to ADD
WORKDIR /app
ADD app.js /app
ADD package.json /app

RUN npm install

CMD ["node", "app.js"]
