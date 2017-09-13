BIN = hello

.PHONY: all clean

all:
	go build -o $(BIN) .

clean:
	rm -f $(BIN)
