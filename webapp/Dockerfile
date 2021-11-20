FROM node:12.18.2

WORKDIR /app/webapp

COPY package.json package.json
RUN npm install

COPY . .
RUN npm run build

CMD ["npm", "start"]

