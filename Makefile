gen-reset:
	@rm ./src/public/index.js ./src/public/global.css ./views/index_templ.go
gen-js:
	@npx tsc -w	
gen-css:
	@npx tailwindcss -i ./src/input.css -o ./src/public/global.css --watch
gen-templ:
	@templ generate
build: gen-templ
	@go build -o bin/site
run: build 
	@./bin/site