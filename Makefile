.PHONY: css build dev clean run

# Build the CSS using Tailwind CLI
css:
	@mkdir -p build/assets/css
	./node_modules/.bin/tailwindcss -i ./styles/globals.css -o ./build/assets/css/style.css --minify

# Build the Go application (generates static files)
build: css
	go run cmd/builder/main.go cmd/builder/definitions.go

# Development mode: watch for changes
dev:
	go run cmd/builder/main.go cmd/builder/definitions.go --dev

# Clean build artifacts
clean:
	rm -rf build assets/css/style.css
