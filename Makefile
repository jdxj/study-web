name = "study-web.out"

build: clean
	go build -o $(name) *.go
clean:
	find . -name "*.log" | xargs rm -f
	find . -name "*.out" | xargs rm -f