# Build the application
all: build

build:
	echo "Building application..."

	go build -o main cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run tests
test:
	echo "Running tests"
	go test ./tests -v

# Clean the binary
clean:
	echo "Cleaning binary.."
	rm -f main
