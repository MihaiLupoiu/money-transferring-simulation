# money-transferring-simulation

# Used:
* Dep
* Gin framework
* SQLite DB

# TODOs:

* Check if mail already in use.
* Add integrity check method to don't do the same transfer several times. (User does several request in one second/minute/day?).
* Change database to a more schalabe one
* Split users and balance
* Add token and SSO as autentication method.
* Change transference method to be done by multiple seperated process and have an organizer to send jobs in order to scale.
* Add Tests