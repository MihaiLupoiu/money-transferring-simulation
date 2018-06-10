# money-transferring-simulation

To build execute:
```bash 
cd ./scripts
./services-build.sh
```

This builds the binary using an alpine and creates the docker image ready to use.


After that there should be a mihailupoiu/users image. To execute run:

```bash 
docker run mihailupoiu/users -p 80:8080
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