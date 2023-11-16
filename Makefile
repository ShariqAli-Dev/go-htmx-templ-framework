gen-css:
	@npx tailwindcss -i ./src/input.css -o ./src/public/global.css --watch
gen-templ:
	@templ generate
build: gen-templ
	@go build -o bin/site
run: build 
	@./bin/site