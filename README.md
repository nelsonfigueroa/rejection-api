# Rejection API

Get the true Software Engineer job search experience.

This is a simple JSON API built on Ruby on Rails. The API returns a collection of rejection emails I have received this year. Messages have been slightly modified to remove company names.

# Running the API

Assuming you have Go installed, run `main.go`:

```
go run main.go
```

You can also download and run the appropriate binary on [the releases section](https://github.com/nelsonfigueroa/rejection_api/releases/tag/v0.1).

Now you can make API calls to [127.0.0.1:3000](http://127.0.0.1:3000).

| HTTP Verb | Endpoint   | Description                    |
|-----------|------------|--------------------------------|
| GET       | /rejections | Get a random rejection message 

**Examples:**

Get a random rejection:

```
curl 127.0.0.1:3000/rejections
```

Response

```json
{"id":5,"message":"We regret to inform you that you are no longer being considered for the Junior Full-Stack Developer position. Thank you for your interest and for taking the time to apply.","created_at":"2020-07-27T23:21:49.043Z","updated_at":"2020-07-27T23:21:49.043Z"}
```
