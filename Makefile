wire-gen:
	wire ./cmd/...

migraion:
	cd migration; \
	goose mysql "test:test@(10.1.1.111:3307)/kbc2?parseTime=true" up;
	# SIT DB migraion:
	# goose mysql "sitkbc:9n#yz635%GntpMfy@(sit-kbc-mysql-tw-01.silkrode.in:5306)/kbc2_captain_marvel?parseTime=true" up

build-captain_marvel:
	./build/build.sh
