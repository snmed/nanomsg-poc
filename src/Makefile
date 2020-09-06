
BIN = nano-cli
INC_DIR = -I./nng/include
LIBS = -lnng
LIB_PATH = -L./nng
CC = clang++
CFLAGS = $(INC_DIR) -std=c++14 -stdlib=libc++ -g -Wall
LDFLAGS = $(LIBS) $(LIB_PATH) 

all: clean bin

bin: *.o
	$(CC) $(CFLAGS) $(LDFLAGS) -o $(BIN) $^

%.o: %.cc
	$(CC) $(CFLAGS) -c $<

clean:
	rm -f *.o
	rm -f $(BIN)

PHONY: all clean bin
