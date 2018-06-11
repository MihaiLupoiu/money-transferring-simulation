# money-transferring-simulation



To build for alpine:
```bash 
dep ensure
cd ./scripts
./services-build.sh
```
To build for local:
```bash 
dep ensure
cd ./services/user
go build .
```

This builds the binary using an alpine and creates the docker image ready to use.


After that there should be a mihailupoiu/users image. To execute run:

```bash 
docker run --rm -p 8080:8080 mihailupoiu/users 
```

# Used:
* Dep
* Gin framework
* SQLite DB

# TODOs:

* Check if mail already in use.
* Add integrity check method to don't do the same transfer several times. (User does several request in one second/minute/day?).
* Change database to a more scalable one
* Split users and balance
* Add token and SSO authentication method.
* Change transference method to be done by multiple separate process and have an organizer to send jobs in order to scale.
* Add Tests