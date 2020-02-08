run:
	@echo ":::: App is startin up ::::"
	@echo "CONFIG::  😁 Exporting environemnt variables"
	# Parrot os source alternative
	/bin/sh .env
	@echo "SUCCESS:  ✔ Environment variables exported"
	@echo "INIT::::  ⚡ Running server"
	go run app.go