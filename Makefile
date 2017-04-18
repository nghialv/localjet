.PHONY: save-vendor
save-vendor:
	rm -rf Godeps/ vendor/
	godep save ./...

.PHONY: compose
compose:
	docker-compose -p compose run --rm testrunner bash -c "sleep 15 && go test -v github.com/nghialv/localjet/mongodb"
	docker-compose -p compose down &>/dev/null &
