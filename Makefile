include .env.dist .env

export SESSION_COOKIE
export DAY

run: create
	go run -mod=readonly ${YEAR}/${DAY}/* --part 2

create:
	@if [ ! -f "${YEAR}/${DAY}/main.go" ]; then \
		mkdir -p ${YEAR}/${DAY}; \
		cp -r template/go/* ${YEAR}/${DAY}/; \
	fi
