format-templ:
	@go run scripts/template_formatter.go
gen-css:
	@npx tailwindcss -i ./web/src/input.css -o ./web/public/global.css --watch
gen-templ:
	@templ generate
build: gen-templ
	@go build -o bin/site
run: build 
	@./bin/site