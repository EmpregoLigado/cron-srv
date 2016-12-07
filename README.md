## Cron Srv

* Cron Scheduler Microservice.

## Setup
With [golang](https://golang.org/doc/install) installed

Install the dependencies listed in `glide.yml` with [glide](https://github.com/Masterminds/glide)
```
glide install
```

Env vars
```bash
export CRON_SRV_DB="postgresql://postgres@localhost/cron_srv_dev?sslmode=disable"
export CRON_SRV_PORT=3000
```
> **Note:** You must have created the database 'cron_srv_dev' in postgres running at localhost (or replace with valid database name and IP);

## Installation
```sh
mkdir -p $GOPATH/src/github.com/EmpregoLigado
cd $GOPATH/src/github.com/EmpregoLigado 
git clone https://github.com/EmpregoLigado/cron-srv.git
cd cron-srv
go build
```
## Running server
```
./cron-srv
# => Starting Cron Service at port 3000
```
### Examples
  - with server running, access the url [localhost:3000/v1/crons](http://localhost:3000/v1/crons) and you should see a empty list `[]`
  - now we are going to insert some crons, execute in terminal:
```
curl -X POST -H "Content-Type: application/json" -d '{"url":"example.com", "expression": "just-test", "status": "active", "max_retries": 2, "retry_timeout": 3}' http://localhost:3000/v1/cron
```
and one new cron will be listed in [localhost:3000/v1/crons](http://localhost:3000/v1/crons)

## API Documentation
|HTTP verb|path|handle||
|:--|:--|:--|:--|
|GET|/v1/healthz|HealthzIndex|retun a state of server `{"alive":true}`|
|GET|/v1/crons|CronIndex|display a list of all crons|
|POST|/v1/cron|CronCreate|create a new cron|
|GET|/v1/cron/:id|CronShow|display a specific cron|
|PUT|/v1/cron/:id|CronUpdate|update a specific cron|
|DELETE|/v1/cron/:id|CronDelete|delete a specific cron|

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges
[![CircleCI](https://circleci.com/gh/EmpregoLigado/cron-srv.svg?style=svg)](https://circleci.com/gh/EmpregoLigado/cron-srv)
[![](https://images.microbadger.com/badges/image/rafaeljesus/cron-srv.svg)](https://microbadger.com/images/rafaeljesus/cron-srv "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/rafaeljesus/cron-srv.svg)](https://microbadger.com/images/rafaeljesus/cron-srv "Get your own version badge on microbadger.com")
