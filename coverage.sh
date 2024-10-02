go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./... ; \
cat coverprofile_.tmp | grep -v mock.go | grep -v _mocks.go | grep -v _gen.go > coverprofile.tmp ; \
rm coverprofile_.tmp ; \
go tool cover -html coverprofile.tmp -o coverage.html ; \
go tool cover -func coverprofile.tmp
