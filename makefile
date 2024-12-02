# Makefile

# Phony target to allow dynamic day arguments
.PHONY: d

# Generic rule to run the script for a given day
d%:
	go run ./src/day$*/main.go
