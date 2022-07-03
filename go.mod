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
