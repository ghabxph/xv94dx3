# xv94dx3 - Covid Statistics

[![Official Website](https://img.shields.io/badge/Official%20Website-https://covid.ghabxph.info-green.svg)](https://covid.ghabxph.info)
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

Then open two terminals. First one:

``` sh
make build up ps log
```

What the above command does is:

1. Builds the xv94dx3 app.
2. Runs xv94dx3, postgresql, and adminer (just a postgresql db explorer like phpmyadmin)
3. Shows running process (ps)
4. Shows docker logs (docker-compsoe logs)

I expect that on process list, the web instance is down. That is because postgresql db
is still down the moment it runs. Now for your second terminal:

``` sh
make build up ps
```

On the first terminal, you will see what's going on to all containers. You must see the 

At first attempt, `xv94dx3_web_1` will die. That is because postgresql was not loaded yet.
Simply rerunning `make build up ps` should rerun it and all data from `covid_19_data.csv`
will be loaded. You must see something like this...

``` 
web_1      | 2021/03/29 13:40:40 &{15107 2020-04-14 00:00:00 +0000 UTC Sint Maarten Netherlands 2020-04-14 23:41:11 +0000 UTC 52 9 5}
web_1      | 2021/03/29 13:40:40 &{15108 2020-04-14 00:00:00 +0000 UTC South Australia Australia 2020-04-14 23:41:11 +0000 UTC 433 4 240}
web_1      | 2021/03/29 13:40:40 &{15109 2020-04-14 00:00:00 +0000 UTC South Carolina US 2020-04-14 23:41:11 +0000 UTC 3553 97 0}
web_1      | 2021/03/29 13:40:40 &{15110 2020-04-14 00:00:00 +0000 UTC South Dakota US 2020-04-14 23:41:11 +0000 UTC 988 6 0}
web_1      | 2021/03/29 13:40:40 &{15111 2020-04-14 00:00:00 +0000 UTC St Martin France 2020-04-14 23:41:11 +0000 UTC 32 2 11}
web_1      | 2021/03/29 13:40:40 &{15112 2020-04-14 00:00:00 +0000 UTC Tasmania Australia 2020-04-14 23:41:11 +0000 UTC 165 6 53}
web_1      | 2021/03/29 13:40:40 &{15113 2020-04-14 00:00:00 +0000 UTC Tennessee US 2020-04-14 23:41:11 +0000 UTC 5827 124 0}
web_1      | 2021/03/29 13:40:40 &{15114 2020-04-14 00:00:00 +0000 UTC Texas US 2020-04-14 23:41:11 +0000 UTC 15006 342 0}
web_1      | 2021/03/29 13:40:40 &{15115 2020-04-14 00:00:00 +0000 UTC Tianjin Mainland China 2020-04-14 23:41:11 +0000 UTC 185 3 168}
web_1      | 2021/03/29 13:40:40 &{15116 2020-04-14 00:00:00 +0000 UTC Tibet Mainland China 2020-04-14 23:41:11 +0000 UTC 1 0 1}
web_1      | 2021/03/29 13:40:40 &{15117 2020-04-14 00:00:00 +0000 UTC Turks and Caicos Islands UK 2020-04-14 23:41:11 +0000 UTC 10 1 0}
web_1      | 2021/03/29 13:40:40 &{15118 2020-04-14 00:00:00 +0000 UTC Utah US 2020-04-14 23:41:11 +0000 UTC 2417 18 0}
web_1      | 2021/03/29 13:40:40 &{15119 2020-04-14 00:00:00 +0000 UTC Vermont US 2020-04-14 23:41:11 +0000 UTC 752 29 0}
web_1      | 2021/03/29 13:40:40 &{15120 2020-04-14 00:00:00 +0000 UTC Victoria Australia 2020-04-14 23:41:11 +0000 UTC 1291 14 1118}
web_1      | 2021/03/29 13:40:40 &{15121 2020-04-14 00:00:00 +0000 UTC Virgin Islands US 2020-04-14 23:41:11 +0000 UTC 51 1 0}
web_1      | 2021/03/29 13:40:40 &{15122 2020-04-14 00:00:00 +0000 UTC Virginia US 2020-04-14 23:41:11 +0000 UTC 6182 154 0}
web_1      | 2021/03/29 13:40:40 &{15123 2020-04-14 00:00:00 +0000 UTC Washington US 2020-04-14 23:41:11 +0000 UTC 10799 530 0}
web_1      | 2021/03/29 13:40:40 &{15124 2020-04-14 00:00:00 +0000 UTC West Virginia US 2020-04-14 23:41:11 +0000 UTC 640 9 0}
web_1      | 2021/03/29 13:40:40 &{15125 2020-04-14 00:00:00 +0000 UTC Western Australia Australia 2020-04-14 23:41:11 +0000 UTC 527 6 251}
web_1      | 2021/03/29 13:40:40 &{15126 2020-04-14 00:00:00 +0000 UTC Wisconsin US 2020-04-14 23:41:11 +0000 UTC 3555 170 0}
web_1      | 2021/03/29 13:40:40 &{15127 2020-04-14 00:00:00 +0000 UTC Wyoming US 2020-04-14 23:41:11 +0000 UTC 282 1 0}
web_1      | 2021/03/29 13:40:40 &{15128 2020-04-14 00:00:00 +0000 UTC Xinjiang Mainland China 2020-04-14 23:41:11 +0000 UTC 76 3 73}
web_1      | 2021/03/29 13:40:40 &{15129 2020-04-14 00:00:00 +0000 UTC Yukon Canada 2020-04-14 23:41:11 +0000 UTC 8 0 0}
web_1      | 2021/03/29 13:40:40 &{15130 2020-04-14 00:00:00 +0000 UTC Yunnan Mainland China 2020-04-14 23:41:11 +0000 UTC 184 2 175}
web_1      | 2021/03/29 13:40:40 &{15131 2020-04-14 00:00:00 +0000 UTC Zhejiang Mainland China 2020-04-14 23:41:11 +0000 UTC 1267 1 1242}
web_1      | 2021/03/29 13:40:40 read through csv done
```

## Accessing the application

As per defined in `docker-compose.yaml`, simply access your local machine to port 8082: `http://localhost:8082`

Swagger shall guide you how to use the api. Of course you can consume the API directly by yourself, just using
swagger as your point of reference.

If you desire to access it on my website, please choose `https` option because my cloudflare is setup only for
`https` connections only. And who uses `http` nowadays? (except local development).

## API Documentation

Our API is documented using Swagger OpenAPI. The documentation is served within  the  app.
The root (/) route shall redirect you to the  API  documentation  and  teach  you  how  to
consume this endpoint.


# Contribution Guide

Feel free to fork this application and make your  modifications.  If  you're  generous  to
share it to us, feel free to open a ticket and send us a pull request.
