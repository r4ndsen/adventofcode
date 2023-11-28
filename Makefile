include .env.dist .env

export COOKIE
export DAY

run:
	@mkdir -p ${YEAR}/${DAY} && \
	touch ${YEAR}/${DAY}/main.go && \
	go run ${YEAR}/${DAY}/main.go

a:
	echo ${COOKIE} ${DAY}
