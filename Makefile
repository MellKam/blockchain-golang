BINARY_NAME=Blockchain
BUILD_PATH=./bin/
DATABASE_FOLDER=./tmp/chain/
EXECUTABLE_FILE_PATH=./cmd/blockchain/main.go

build: 
	make build_macos
	make build_linux
	make build_windows

build_macos:
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_PATH}${BINARY_NAME}-macos ${EXECUTABLE_FILE_PATH}

build_linux:
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_PATH}${BINARY_NAME}-linux ${EXECUTABLE_FILE_PATH}

build_windows:
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_PATH}${BINARY_NAME}-windows ${EXECUTABLE_FILE_PATH}

clean_build:
	rm -r ${BUILD_PATH}

clean_db:
	rm -r ${DATABASE_FOLDER}

clean:
	make clean_db
	make clean_build
	
run:
	go run ${EXECUTABLE_FILE_PATH} $(ARGS)