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

## Building and running the application

Clone our repository:

``` sh
git clone https://github.com/ghabxph/xv94dx3.git
```

Then by using make, conveniently build and turn the docker containers up:

``` sh
make build up ps
```

`ps` is just for showing the process.

If `xv94dx3_web_1` has exited, it means that postgres was not  loaded  up  the  time  that
xv94dx3 is loading up. To fix that, simply run `make build up ps` again and it should  run
by now.

## API Documentation

Our API is documented using Swagger OpenAPI. The documentation is served within  the  app.
The root (/) route shall redirect you to the  API  documentation  and  teach  you  how  to
consume this endpoint.

# Contribution Guide

Feel free to fork this application and make your  modifications.  If  you're  generous  to
share it to us, feel free to open a ticket and send us a pull request.
