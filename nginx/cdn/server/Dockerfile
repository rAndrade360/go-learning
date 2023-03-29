FROM nginx:latest

RUN rm -rf /etc/nginx/conf.d/*

COPY server.conf /etc/nginx/conf.d/

RUN mkdir /spool && mkdir /spool/logs && touch /spool/logs/lb-access.log && mkdir /spool/dev && mkdir /spool/dev/performance

COPY *.html /spool/dev/performance/

COPY *.png /spool/dev/performance/
