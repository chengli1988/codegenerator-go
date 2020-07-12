module codegenerator-go

go 1.12

require (
	codegenerator-go/models v0.0.0
	codegenerator-go/utils v0.0.0
	github.com/chengli1988/dbutil-mysql v0.0.0-20200711164911-ce7ffc0f4e19
)

replace codegenerator-go/models => ./models

replace codegenerator-go/utils => ./utils
