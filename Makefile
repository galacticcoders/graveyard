.PHONY: validate

validate:
	cd tools/validate-graves && go run . ../../graves.json
