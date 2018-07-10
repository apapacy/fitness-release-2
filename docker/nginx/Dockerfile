FROM openresty/openresty:alpine

RUN apk update \
 && apk add supervisor \
 && apk add logrotate \
 && apk add openssl \
 && openssl dhparam -dsaparam -out /etc/ssl/certs/dhparam.pem 4096

COPY ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY ./logrotate.conf /etc/logrotate.conf
COPY ./cron.d /etc/cron.d/nginx

ARG UID

RUN mkdir -p /var/log/supervisor \
  && chmod 644 /etc/logrotate.conf && chown root:root /etc/logrotate.conf \
  && chmod 644 /etc/cron.d/nginx && chown root:root /etc/cron.d/nginx \
  && addgroup -S app && adduser -S -G app app -u $UID

ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"] 
