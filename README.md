# Rejection API

Simple API built on Ruby on Rails. This is a collection of rejection emails I have received this year. Messages have been slightly modified to remove company names.

# Try it yourself

Pull the Docker image

```
docker pull nfigueroa/rejection-api
```

Run the Docker image

```
docker run -p 3000:3000 nfigueroa/rejection-api
```

Make API calls to [127.0.0.1:3000/rejection](http://127.0.0.1:3000/rejection). There are 2 endpoints.

Get a random rejection message:

- GET /rejection

Example:

```
curl 127.0.0.1:3000/rejection
```

Create a rejection message:

- POST /rejection

Example:

```
curl -d 'message=Sorry college did not prepare you, but we do not want to train anyone. Good luck!' -X POST 127.0.0.1:3000/rejection
```
