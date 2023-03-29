FROM nginx:latest

RUN rm -rf /etc/nginx/conf.d/*

COPY cdn.conf /etc/nginx/conf.d/

RUN mkdir /spool && mkdir /spool/logs && touch /spool/logs/lb-access.log && mkdir /spool/dev && mkdir /spool/dev/cdn

COPY *.png /spool/dev/cdn/
