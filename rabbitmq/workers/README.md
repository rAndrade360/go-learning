# worker-api

To start the api, run:
```bash
$ docker compose up
```

To curl the api, run:
```bash
$ curl -iv -d '{"message": "ble"}' 'http://0.0.0.0:8080/pub'
```