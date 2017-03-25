## Cron Srv

* All the flexibility and power of Cron as a Service.
* Simple REST protocol, integrating with a web application in a easy and straightforward way.
* No more wasting time building and managing scheduling infrastructure.

## Basic Concepts
Cron Srv works by calling back to your application via HTTP GET according to a schedule constructed by you or your application.

## Setup
Env vars
```bash
export CRON_SRV_DB="postgresql://postgres@localhost/cron_srv_dev?sslmode=disable"
export CRON_SRV_PORT=3000
```
> **Note:** You must have created the database 'cron_srv_dev' in postgres running at localhost (or replace with valid database name and IP);

```sh
mkdir -p $GOPATH/src/github.com/EmpregoLigado
cd $GOPATH/src/github.com/EmpregoLigado 
git clone https://github.com/EmpregoLigado/cron-srv.git
cd cron-srv
glide install
go build
```

## Running server
```
./cron-srv
# => Starting Cron Service at port 3000
```

### Create an Cron
- Request
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "example.com/api/v1/stats", "expression": "0 5 * * * *", "status": "active", "max_retries": 2, "retry_timeout": 3}' \
localhost:3000/v1/events
```

- Response
```json
{
  "id": 1,
  "url": "example.com/api/v1/stats",
  "expression": "0 5 * * * *",
  "status": "active",
  "max_retries": 2,
  "retry_timeout": 3,
  "updated_at": "2016-12-10T14:02:37.064641296-02:00"
}
```

## API Documentation
|HTTP verb| path|                   handle|
|:--|:--|:--|:--|
|GET        |/v1/healthz|HealthzIndex     |return a state of server `{"alive":"up"}`|
|GET        |/v1/events|EventsIndex       |display a list of all events|
|POST       |/v1/events|EventsCreate      |create a new event|
|GET        |/v1/events/:id|EventsShow    |display a specific event|
|PUT        |/v1/events/:id|EventsUpdate  |update a specific event|
|DELETE     |/v1/events/:id|EventsDelete  |delete a specific event|

## Cron Format
The cron expression format allowed is:

|Field name| Mandatory?|Allowed values|Allowed special characters|
|:--|:--|:--|:--|
|Seconds      | Yes        | 0-59            | * / , -|
|Minutes      | Yes        | 0-59            | * / , -|
|Hours        | Yes        | 0-23            | * / , -|
|Day of month | Yes        | 1-31            | * / , - ?|
|Month        | Yes        | 1-12 or JAN-DEC | * / , -|
|Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?|
more details about expression format [here](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format)

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges
[![CircleCI](https://circleci.com/gh/EmpregoLigado/cron-srv.svg?style=svg)](https://circleci.com/gh/EmpregoLigado/cron-srv)
[![Go Report Card](https://goreportcard.com/badge/github.com/EmpregoLigado/cron-srv)](https://goreportcard.com/report/github.com/EmpregoLigado/cron-srv)
[![](https://images.microbadger.com/badges/image/rafaeljesus/cron-srv.svg)](https://microbadger.com/images/rafaeljesus/cron-srv "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/rafaeljesus/cron-srv.svg)](https://microbadger.com/images/rafaeljesus/cron-srv "Get your own version badge on microbadger.com")
