module codegenerator-go

go 1.12

require (
	codegenerator-go/models v0.0.0
	codegenerator-go/utils v0.0.0
	github.com/chengli1988/go-dbutil-mysql v0.0.0-20191029064344-dbaa6e4ada2b
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
)

replace codegenerator-go/models => ./models

replace codegenerator-go/utils => ./utils
