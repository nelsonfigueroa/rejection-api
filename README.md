# Rejection API

Get the true Software Engineer job search experience.

This is a simple JSON API built on Ruby on Rails. The API returns a collection of rejection emails I have received this year. Messages have been slightly modified to remove company names.

# Running the API

Pull the Docker image

```
docker pull nfigueroa/rejection-api
```

Run the Docker image

```
docker run -p 3000:3000 nfigueroa/rejection-api
```

Now you can make API calls to [127.0.0.1:3000/rejection](http://127.0.0.1:3000/rejection). There are 2 endpoints.

| HTTP Verb | Endpoint   | Description                    |
|-----------|------------|--------------------------------|
| GET       | /rejections | Get a random rejection message |
| POST      | /rejections | Create a rejection message     |

**Examples:**

Get a random rejection
```
curl 127.0.0.1:3000/rejections
```

Response

```json
{"id":20,"message":"We want to thank you for taking the time to complete our assessment regarding our Software Engineer position.  After careful review it has been decided that there is not a strong fit at the moment.   Thanks again for your interest in us and I hope you will remain enthusiastic about our company.","created_at":"2020-07-21T08:18:31.604Z","updated_at":"2020-07-21T08:18:31.604Z"}
```

Create a rejection
```
curl -d 'message=Sorry college did not prepare you, but we do not want to train anyone. Good luck!' -X POST 127.0.0.1:3000/rejections
```

Response

```json
{"id":26,"message":"Sorry college did not prepare you, but we do not want to train anyone. Good luck!","created_at":"2020-07-21T18:13:26.679Z","updated_at":"2020-07-21T18:13:26.679Z"}
```