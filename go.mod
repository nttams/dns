module main

go 1.18

replace message => ./message

replace client => ./client

replace server => ./server

replace zone_handler => ./zone_handler

require (
	client v0.0.0-00010101000000-000000000000
	message v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
	zone_handler v0.0.0-00010101000000-000000000000
)

require (
	github.com/k0kubun/pp v3.0.1+incompatible // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
)
