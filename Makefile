include .env.dist .env

export SESSION_COOKIE
export DAY

run:
	@mkdir -p ${YEAR}/${DAY} && \
    touch ${YEAR}/${DAY}/main.go && \
	go run -mod=readonly ${YEAR}/${DAY}/* --part 0

