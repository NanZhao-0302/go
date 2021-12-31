module learn.go

go 1.17

require github.com/spf13/cobra v1.3.0

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
