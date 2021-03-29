# xv94dx3 - Covid Statistics

[![Maintainer - GhabXPH](https://img.shields.io/badge/maintainer-ghabxph-green.svg)](https://github.com/ghabxph)
[![Project Status - Feature Complete](https://img.shields.io/badge/project%20status-feature%20complete-green.svg)](https://github.com/ghabxph/xv94dx3)

Just a simple tool that returns a list of the top N countries with confirmed cases  for  a
given observation date where N is the maximum number of results. Include the total  number
of deaths and recoveries per country in the response for that day.

What's behind the name xv94dx3? Well it may be one of my passwords because I just generated
this name on a password manager. LOL.

## Requirements

* GNU Make (optional)
* Docker
* Docker Compose

## Pre-setup

Clone our repository:

``` sh
git clone https://github.com/ghabxph/xv94dx3.git
```

Then go to `docker-compose.yaml`. Uncomment  `#command: --runSeeder=covid_19_data.csv`  by
removing the hash symbol. This is to import all data from our  csv file  to  our  postgres
database.

Then build and run our app by running:

``` sh
make build up ps
```

At first attempt, `xv94dx3_web_1` will die. That is because postgresql was not loaded yet.
Simply rerunning `make build up ps` should rerun it and all data from `covid_19_data.csv`
will be loaded.

## Building and running the application

By using make, conveniently build and turn the docker containers up:

``` sh
make build up ps
```

`ps` is just for showing the process.

## API Documentation

Our API is documented using Swagger OpenAPI. The documentation is served within  the  app.
The root (/) route shall redirect you to the  API  documentation  and  teach  you  how  to
consume this endpoint.

# Contribution Guide

Feel free to fork this application and make your  modifications.  If  you're  generous  to
share it to us, feel free to open a ticket and send us a pull request.
