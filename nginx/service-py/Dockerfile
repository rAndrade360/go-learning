FROM python:latest

RUN apt-get update; apt-get install -y curl

COPY . .

CMD [ "python3", "server.py" ]
