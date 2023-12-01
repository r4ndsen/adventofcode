include .env.dist .env

export SESSION_COOKIE
export DAY

run:
	@mkdir -p ${YEAR}/${DAY} && \
	go run -mod=readonly ${YEAR}/${DAY}/main.go --part 2
