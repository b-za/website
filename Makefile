.PHONY: css build dev clean run

# Build the CSS using Tailwind CLI
css:
	./node_modules/.bin/tailwindcss -i ./styles/globals.css -o ./assets/css/style.css --minify

# Build the Go application (generates static files)
build: css
	go run cmd/builder/main.go cmd/builder/definitions.go

# Development mode: watch for changes
dev:
	@echo "Starting development server..."
	@# distinct simple watcher or just instructions to use air if installed
	go run cmd/builder/main.go definitions.go --dev

# Clean build artifacts
clean:
	rm -rf build assets/css/style.css
