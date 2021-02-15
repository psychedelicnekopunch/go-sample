# https://hub.docker.com/_/nginx
FROM nginx:1

WORKDIR /var/www/html

COPY ./app /var/www/html
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
