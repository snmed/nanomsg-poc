
BIN = nano-cli
LIBS = -lnng
CC = clang++
BASE_DIR = $(shell pwd)
BUILD_DIR = $(BASE_DIR)/nng/build
INC_DIR = -I$(BASE_DIR)/nng/include
LIB_PATH = -L$(BUILD_DIR)
LDFLAGS = $(LIBS) $(LIB_PATH)
CFLAGS = $(INC_DIR) -std=c++14 -stdlib=libc++ -g -Wall

.PHONY: all
all: clean nng bin

.PHONY: nng
nng:
	mkdir -p $(BUILD_DIR)
	cd $(BUILD_DIR) && cmake ..
	make -C $(BUILD_DIR) nng

.PHONY: bin
bin: *.o	
	$(CC) $(CFLAGS) $(LDFLAGS) -o $(BIN) $^

%.o: %.cpp
	$(CC) $(CFLAGS) -c $<

.PHONY: clean
clean:
	rm -f *.o
	rm -f $(BIN)
	rm -rf $(BUILD_DIR)


