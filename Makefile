format-templ:
	@go run scripts/template_fromatter.go
gen-js:
	@npx tsc -w
gen-css:
	@npx tailwindcss -i ./web/input.css -o ./web/public/global.css --watch
gen-templ:
	@templ generate
build: gen-templ
	@go build -o bin/site
run: build 
	@./bin/site