# Этап сборки на Node.js
FROM node:18-alpine AS build

WORKDIR /app

# Копируем package.json и package-lock.json для кеширования зависимостей
COPY package.json package-lock.json ./
RUN npm ci

# Копируем остальной код и собираем фронтенд
COPY . . 
RUN npm run build

# Этап развертывания в Nginx
FROM nginx:stable-alpine
COPY --from=build /app/dist /usr/share/nginx/html

# Настраиваем Nginx
COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
