make:
	if ! [ -x "$$(command -v packr)" ]; then \
		go get -u github.com/gobuffalo/packr/packr; \
		packr; \
	fi
	go install