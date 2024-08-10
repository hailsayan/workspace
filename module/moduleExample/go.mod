module github.com/hailsayan/moduleExample

go 1.18

require github.com/jalaali/go-jalaali v0.0.0-20210801064154-80525e88d958

require github.com/hailsayan/TestModule v0.0.1 // indirect

replace github.com/jalaali/go-jalaali => ./local/go-jalaali
