FROM nginx:latest

COPY lb.conf /etc/nginx/conf.d/

RUN mkdir /spool && mkdir /spool/logs && touch lb-access.log

RUN nginx -t
