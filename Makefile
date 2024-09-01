dev:
	make -j3 dev/server dev/templ dev/tailwind

dev/server:
	air \
	--build.cmd "go build -o tmp/main cmd/main/main.go" --build.bin "tmp/main" --build.delay "100" \
	--build.exclude_dir "web/node_modules" \
	--build.include_ext "go"

dev/templ:
	templ generate --watch --open-browser=false

dev/tailwind:
	cd web && npm run dev:css
