# fitness-release-2

go get -u github.com/gin-gonic/gin/...

go get -u gopkg.in/src-d/go-kallax.v1/...

go get github.com/oklog/ulid

kallax migrate --input ./model/ --out ./migrations --name init_db


kallax migrate up --dir ./migrations --dsn 'postgres:password@localhost:5433/test?sslmode=disable' --steps 1
kallax migrate up --dir ./my-migrations --dsn 'user:pass@localhost:5432/dbname?sslmode=disable' --version 1493991142

go generate ./model/...

export GOROOT=$HOME/go
export GOPATH=$HOME/projects/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export GOCACHE=off go test
go test -v model/*


go get -u github.com/golang/dep/cmd/dep
go get -u -d github.com/golang-migrate/migrate/cli
cd $GOPATH/src/github.com/golang-migrate/migrate/cli
dep ensure
go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cli

 migrate -database postgres://postgres:password@localhost:5433/test?sslmode=disable  -path migrations drop
