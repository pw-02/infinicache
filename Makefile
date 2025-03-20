deploy:
	aws configure
	
build:
	go build -o proxy/proxy github.com/mason-leap-lab/infinicache/proxy

start: build
	proxy/proxy 1>./log 2>&1 &

stop:
	@if pgrep -f "proxy/proxy" > /dev/null; then \
		kill -2 $$(pgrep -f "proxy/proxy") && echo "proxy/proxy process stopped."; \
	else \
		echo "No proxy/proxy process found."; \
	fi
