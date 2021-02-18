** NF_STN **

REST API developed in Go with the following features:

- List invoices with invoices

- Search invoice by document

- Register invoice

- Logical invoice deletion

- List invoices with filters by month, year and document

- List invoices sorted by month, year, document or combinations of these

Persistence developed using the MySQL database.

Route authentication by application token.

** How to run the code **

To run the application just run the following command on your terminal:

docker-compose up -d

Running the tests:

go test. / ...

** How the API works **

To access the endpoints it is necessary to login the user, the bank
data is already started with a user table that has a row
containing a "username" and a "hash" generated from the password "password".
The application then compares the hashes and authorizes the user. Case
the user is authenticated, the application returns a token of the type "Bearer Token" to be used
in the authentication header for the other routes. Such token is valid for 15 minutes
can be renewed at the endpoint api / refresh, and is stored on the Redis server for
future checks.

** Search parameters **

For ordering use the parameter "orderBy" in the request, with the following ordering options:
"document", "year," month "or a composition separated by commas.

To search by month, year or document use the parameters "month", "year"
or "document" respectively.

Following is a postman collection with examples of requests (use
a new token).

** Deploy the application on AWS **

The application was deployed on AWS and to access it from any machine
just access the following address: http://18.188.115.62:8000 (instance not up anymore)

Remembering that to log in you need to add in the request header of the endpoint / api / login
the parameters "username" = username and "password" = password.
