BINARY_NAME=Blockchain
BUILD_PATH=./bin/
DATABASE_FOLDER=./tmp/chain/
EXECUTABLE_FILE_PATH=./cmd/blockchain/main.go

build: 
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_PATH}${BINARY_NAME}-darwin ${EXECUTABLE_FILE_PATH}
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_PATH}${BINARY_NAME}-linux ${EXECUTABLE_FILE_PATH}
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_PATH}${BINARY_NAME}-windows ${EXECUTABLE_FILE_PATH}

clean_build:
	rm -r ${BUILD_PATH}

clean_db:
	rm -r ${DATABASE_FOLDER}
	
run:
	go run ${EXECUTABLE_FILE_PATH}