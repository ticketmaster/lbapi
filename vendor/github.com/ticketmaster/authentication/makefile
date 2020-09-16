gpath = $(subst ${GOPATH}/src/,,${PWD})
all: clean prepare coverage
clean:
	if [ -d "cover" ]; then rm -r cover/; fi
prepare:
	glide install
coverage:
	if [ ! -d "cover" ]; then mkdir cover; fi
	go test -coverprofile cover/cover-authentication.out -covermode count
	go test -coverprofile cover/cover-client.out -covermode count "${gpath}/client"
	go test -coverprofile cover/cover-authorization.out -covermode count  "${gpath}/authorization"
	go test -coverprofile cover/cover-common.out -covermode count  "${gpath}/common"
	cat  cover/cover-authentication.out >> cover/cover.out
	cat  cover/cover-client.out | tail -n +2 >> cover/cover.out
	cat  cover/cover-authorization.out | tail -n +2 >> cover/cover.out
	cat  cover/cover-common.out | tail -n +2 >> cover/cover.out
	go tool cover -html=cover/cover.out -o cover/coverage.html
	go tool cover -func=cover/cover.out
.PHONY: all