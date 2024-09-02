dev:
	make -j3 dev/server dev/templ dev/tailwind

dev_docker:
	make -j3 dev/server_docker dev/templ dev/tailwind

dev/server:
	air

dev/server_docker:
	docker-compose -f compose.dev.yml up --watch

dev/templ:
	templ generate --watch --open-browser=false

dev/tailwind:
	cd web && npm run dev:css
