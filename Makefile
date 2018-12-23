BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

vendor:
	@dep ensure -v

engine: vendor
	go build -o ${BINARY} app/*.go

install: 
	go build -o ${BINARY} app/*.go

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

run:
	docker-compose up -d

stop:
	docker-compose down

.PHONY: clean install unittest build docker run stop vendor