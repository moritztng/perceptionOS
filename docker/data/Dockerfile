FROM node:19

WORKDIR /app

COPY package*.json ./

COPY prisma ./prisma

RUN npm install && npx prisma migrate dev --name init

COPY . .

CMD [ "npm", "start" ]
