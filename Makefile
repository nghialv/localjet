.PHONY: save-vendor
save-vendor:
	rm -rf Godeps/ vendor/
	godep save ./...
