all: gBTRx

gBTRx:
	go build -o bin/gBTRx

clean:
	rm -rf bin/gBTRx