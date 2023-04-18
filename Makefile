PACKAGE_LIST := $(shell go list ./...)
sleet:
        go build -o sleet $(PACKAGE_LIST)
test:
        go test $(PACKAGE_LIST)
clean:
        rm -f sleet
