# Nginx

> This project is a full example of how to use nginx to create a load balancer


## To test

### Requirements

 - Docker
 - cUrl

### Testing

Run the follwing command in your terminal:

```bash
$ docker compose up
```
The containers will be started. Now, run this another command:

```bash
$  curl localhost:8080/
```
The response body may alternate between python and go server response.