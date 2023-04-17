# Имя файла для исполняемого файла вашей функции
BINARY_NAME = main

# Имя файла ZIP-архива, который будет загружен на AWS Lambda
ZIP_NAME = main.zip

all: clean build zip

build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) main.go
	# go build -o $(BINARY_NAME) main.go

zip:
	zip $(ZIP_NAME) $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME) $(ZIP_NAME)
