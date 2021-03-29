bin/xv94dx3:
	@echo ---- building the binary ----
	go build -o bin/xv94dx3 cmd/xv94dx3/main.go

clean:
	@echo ---- cleans the binary folder ----
	rm -fvr bin

run:
	@echo ---- running the app ----
	bin/xv94dx3

run-with-seeder:
	@echo ---- running the app ----
	bin/xv94dx3 --runSeeder=covid_19_data.csv

pull:
	@echo ---- pull docker images ----
	docker-compose pull

build:
	@echo ---- builds images ----
	docker-compose build

up:
	@echo ---- runs docker setup ----
	docker-compose up -d

down:
	@echo ---- stops docker containers ----
	docker-compose down

ps:
	@echo ---- shows docker processes ----
	docker-compose ps

log:
	@echo ---- shows logs -----
	docker-compose logs -f
