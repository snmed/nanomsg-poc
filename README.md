# NNG PoC


## Build Go Client / Server
Build go client:
```bash
# Change into directory go-nano/go-cli
$ go build
```

Build go server:
```bash
# Change into directory go-nano/go-serv
$ go build
```

## Build C++ Client
**Requirments**:
- cmake (Windows & MacOS)
- make (MacOS)

Install submodule with:
```bash
$ git submodule init
$ git submodule update --recursive
```

Switching into the `src` directory and run the following commands:

Windows:
```bash
$ build.bat
```

MacOS:
```bash
$ make
```
