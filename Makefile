.PHONY: copyEnv startDevServer startDevEnvironment run

copyEnv:
	if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi

startDevServer: copyEnv
	air -c .air.toml

startDevEnvironment:
	docker-compose -f .devcontainer/docker-compose.yml up -d

run: copyEnv
	docker-compose up -d
