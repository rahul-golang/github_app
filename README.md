# Github App


This is an example node application that implements Githubs OAuth2 API.

In order to run the application:

1. Register your new application on Github : https://github.com/settings/applications/new. In the "callback URL" field, enter "http://localhost:8080/oauth/redirect". Once you register, you will get a client ID and client secret.
2. Replace the values of the `clientID` and `clientSecret` env variables in the [docker_compose.go](/.go) file and also the [index.html] public/index.html file 
4. Start the server by executing `docker-compose up`
5. Navigate to http://localhost:8080 on your browser.
6. export postman collection 
