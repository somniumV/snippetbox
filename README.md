## Snippetbox - Let's Go

Snippetbox is a full-stack Go application that allows users to share text snippets over CRUD operations, similarly to Pastebin or GitHub's Gist. 
This repository is based on the book (Let's Go by Alex Edwards)[https://lets-go.alexedwards.net/]

### Features
- RESTful Routing
- Authentication 
- CRSF protection
- Middleware
- Database Connection
- Templating
- Testing (Unit, E2E, Integration)

### Routes

| Method | Pattern | Action |
| ------ | ------- | ------ |
| GET | / | Display the home page |
| GET | /snippet/view/:id	| Display a specific snippet |
| GET | /snippet/create	| Display a HTML form for creating a new snippet |
| POST | /snippet/create | Create a new snippet |
| GET | /user/signup | Display a HTML form for signing up a new user |
| POST | /user/signup | Create a new user |
| GET | /user/login | Display a HTML form for logging in a user |
| POST | /user/login | Authenticate and login the user |
| GET | /user/logout | Logout the user |
| GET | /static/*filepath | Serve a specific static file |

### Setup

The local webserver can be starting with the following command  
```
go run ./cmd/web/ 
(-dsn [datasource db string])
(-addr [web port])
```