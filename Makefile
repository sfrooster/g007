PACKAGES := \
	github.com/eliasson/foo \
	github.com/eliasson/bar

DEPENDENCIES := \
	github.com/confluentinc/confluent-kafka-go/kafka \
	github.com/nats-io/go-nats

# all: build silent-test

all: build-k build-n

build-k:
	go build -o bin/g007-k cmd/g007-k/g007-k.go

build-n:
	go build -o bin/g007-n cmd/g007-n/g007-n.go

#test:
#	go test -v $(PACKAGES)

#silent-test:
#	go test $(PACKAGES)

#format:
#	go fmt $(PACKAGES)

#deps:
#	go get $(DEPENDENCIES)

deps:
	dep ensure -add $(DEPENDENCIES)

dep-init:
	dep init

dep-all: dep-init deps