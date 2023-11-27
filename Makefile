format-templ:
	@bin/template_formatter.exe
gen-css:
	@npx tailwindcss -i ./web/src/input.css -o ./web/public/global.css --watch
gen-templ:
	@templ generate
build: gen-templ
	@go build -o bin/site
run: build 
	@./bin/site