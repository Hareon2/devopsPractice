# syntax=docker/dockerfile:1

# 1. For build React app
FROM node:12.14-alpine AS build

# Set working directory
WORKDIR /app

# Install app dependencies
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

#npm install
RUN npm ci

# Copy app source code to the working directory
COPY . /app


ARG REACT_APP_API_URL=""
ENV REACT_APP_API_URL=$REACT_APP_API_URL

# Build stage
RUN npm run build

# 2. For Nginx setup
FROM nginx:alpine

# Copy config nginx
COPY --from=build /app/.nginx/nginx.conf /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

# Remove default nginx static assets
RUN rm -rf ./*

# Copy static assets from builder stage
COPY --from=build /app/build .

# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]