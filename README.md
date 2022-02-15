# Rejection API

Get the true Software Engineer job search experience.

This is a simple JSON API created with Go. The API returns a collection of rejection emails I received in 2020. Messages have been slightly modified to remove company names.

# Running the API

Assuming you have Go installed, get the necessary packages:

```
go get
```

Build the binary:

```
go build
```

Then run it:

```
./rejection-api
```

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
{"id":"2","message":"Thank you for giving us the opportunity to consider you for employment. We have reviewed your background and qualifications and find that we do not have an appropriate position for you at this time."}
```
