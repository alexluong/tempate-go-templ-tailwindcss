dev:
	make -j3 dev/server dev/templ dev/tailwind

dev/server:
	air

dev/templ:
	templ generate --watch --open-browser=false

dev/tailwind:
	cd web && npm run dev:css
