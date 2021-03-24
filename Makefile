bin/xv94dx3:
	@echo ---- building the binary ----
	go build -o bin/xv94dx3 cmd/xv94dx3/main.go

clean:
	@echo ---- cleans the binary folder ----
	rm -fvr bin

run:
	@echo ---- running the app ----
	bin/xv94dx3
