module github.com/rockcookies/go-gen/tests

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/rockcookies/go-gen v0.3.19
	gorm.io/driver/mysql v1.5.7
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.26.0
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/dbresolver v1.6.2
)

replace github.com/rockcookies/go-gen => ../
