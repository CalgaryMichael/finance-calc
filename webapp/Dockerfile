FROM node:12.18.2

WORKDIR /app/webapp
COPY . .

RUN npm install
RUN npm run build

CMD ["npm", "start"]

