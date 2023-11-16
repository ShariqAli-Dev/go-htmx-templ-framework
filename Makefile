gen-css:
	@pnpm dlx tailwindcss -i ./views/input.css -o ./public/global.css --watch
gen-templ:
	@templ generate
build: gen-templ
	@go build -o bin/site
run: build 
	@./bin/site