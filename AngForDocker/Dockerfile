# node base image
FROM node:12.2.0-alpine as build

# working directory
WORKDIR /app

# copy everything to current working directory 
COPY . ./

# run npm install
RUN npm install &&\
  npm run build

# nginx base image
FROM nginx:1.16.0-alpine

# copy static contents of project to nginx html 
COPY --from=build /app/dist/AngForDocker /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]


# Stage 1
#FROM node:10-alpine as build-step
#RUN mkdir -p /app
#WORKDIR /app
#COPY package.json /app
#RUN npm install
#COPY . /app
#RUN npm run build --prod

# Stage 2
#FROM nginx:1.17.1-alpine
#COPY --from=build-step /app/docs /usr/share/nginx/html


