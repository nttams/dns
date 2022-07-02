module main

go 1.18

replace message => ./message
replace zone_handler => ./zone_handler

require (
	message v0.0.0-00010101000000-000000000000 // indirect
	zone_handler v0.0.0-00010101000000-000000000000 // indirect
)
