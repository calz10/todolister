GO=go
GORUN=${GO} run
GOBUILD=${GO} build
GOTEST=${GO} test

CMD_PATH=./cmd/todolister
CMD_FILE=todolister.go

BUILD_PATH=./build
BUILD_NAME=main
BUILD_FILE=$(BUILD_PATH)/$(BUILD_NAME)

RUN=$(GORUN) $(CMD_PATH)/$(CMD_FILE)
BUILD=$(GOBUILD) -o $(BUILD_FILE) $(CMD_PATH)/$(CMD_FILE)


# Environment variables for build mode
DEV_MODE=ENV_MODE=development
PROD_MODE=ENV_MODE=production


# Builds
build-app:
	$(BUILD)
run-dev:
	$(DEV_MODE) $(RUN)
run-prod:
	$(PROD_MODE) $(RUN)
run-build:
	$(PROD_MODE) $(BUILD_FILE)
