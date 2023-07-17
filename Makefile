PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.0
NAME := sleet
DIST := $(NAME)-$(VERSION)
USER_NAME := kimura-hytaae
REPO_NAME := $(USER_NAME)/$(NAME)

$(NAME): coverage.out
	go build -o $(NAME) cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go
#	go install cmd/sleet/main.go

coverage.out: 
	go test -covermode=count \
		-coverprofile=coverage.out $(PACKAGE_LIST)

docker: sleet
#	docker build -t ghcr.io/$(REPO_NAME):$(VERSION) -t ghcr.io/$(REPO_NAME):latest .
	docker buildx build -t ghcr.io/$(REPO_NAME):$(VERSION) \
		-t ghcr.io/$(REPO_NAME):latest --platform=linux/arm64/v8,linux/amd64 --push .

# refer from https://pod.hatenablog.com/entry/2017/06/13/150342
define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
#	cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: $(NAME)
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,386,.exe)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,linux,386,)

distclean : clean
	rm -rf dist

clean :
	rm -f $(NAME) coverage.out
	rm -rf completion cmd/sleet/completions