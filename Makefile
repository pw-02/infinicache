deploy:
	aws configure
	
build:
	go build -o proxy/proxy github.com/mason-leap-lab/infinicache/proxy

start: build
	./proxy/proxy 1>./log 2>&1 & echo $$! > /tmp/proxy.pid

stop:
	@if [ -f /tmp/proxy.pid ]; then \
		PID=$$(cat /tmp/proxy.pid); \
		if ps -p $$PID > /dev/null; then \
			kill -2 $$PID && echo "proxy/proxy process stopped."; \
			rm -f /tmp/proxy.pid; \
		else \
			echo "No running process found with PID $$PID."; \
			rm -f /tmp/proxy.pid; \
		fi \
	else \
		echo "PID file not found, trying pgrep..."; \
		PIDS=$$(pgrep -f "proxy/proxy"); \
		if [ -n "$$PIDS" ]; then \
			kill -2 $$PIDS && echo "proxy/proxy process stopped."; \
		else \
			echo "No proxy/proxy process found."; \
		fi \
	fi
