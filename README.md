# Rejection API

Get the true Software Engineer job search experience.

This is a simple JSON API created with Go. The API returns a selection of rejection emails I've received throughout my career. Messages have been slightly modified to remove company names.

## Running the API by Building the Binary

Make sure you have Go installed. Clone this repo, `cd` into the directory, then get the necessary packages:

```sh
go get
```

Build the binary:

```sh
go build
```

Then run it:

```sh
./rejection-api
```

## Running with Docker

Make sure Docker is running on your system. Clone this repo, `cd` into the directory, then run:

```sh
docker build . -t rejection-api
```

After the image is built, run it with:

```sh
docker run -p 80:80 rejection-api:latest
```

## Using the API

You can make API calls to [127.0.0.1:80](http://127.0.0.1:80).

| HTTP Verb | Endpoint   | Description                    |
|-----------|------------|--------------------------------|
| GET       | /rejections | Get a random rejection message

**Examples:**

Get a random rejection:

```sh
curl -s 127.0.0.1:80/rejections
```

Response

```json
{"id":"2","message":"Thank you for giving us the opportunity to consider you for employment. We have reviewed your background and qualifications and find that we do not have an appropriate position for you at this time."}
```
