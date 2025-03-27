# Run templ generation in watch mode
templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v

# Run air for Go hot reload
server:
	air \
	--build.cmd "go build -o tmp/bin/main.exe ./cmd/server/main.go" \
	--build.bin "tmp\\bin\\main.exe" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# Watch Tailwind CSS changes
tailwind:
	tailwindcss -i ./public/css/input.css -o ./public/css/output.css --watch

# Start development server with all watchers
dev:
	make -j3 templ server tailwind
