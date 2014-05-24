PROFILE_FILE=profile.out
COVERALLS_TOKEN=VJzulX22jziy0V7WHoJpBfA999cMysKCm
TMP_COVER_DIR=tmp_cover

build:
	godep go build -v

test:
	godep go test -race -v `go list ./...` 

run: build
	./godot

cover:
	set -x; \
	mkdir -p ${TMP_COVER_DIR}; \
	for pkg in `go list ./...`; do \
		echo $$pkg; \
		godep go test -v $$pkg -coverprofile=$$(mktemp -p ${TMP_COVER_DIR} -t coverXXX.out) || exit 1; \
	done; \
	echo "mode: set" > ${PROFILE_FILE}; \
	cat ${TMP_COVER_DIR}/cover*.out | grep -v "mode: set" >> ${PROFILE_FILE} || exit 1; \
	rm -rf ${TMP_COVER_DIR}

coveralls: cover
	goveralls -coverprofile=${PROFILE_FILE} -repotoken=${COVERALLS_TOKEN} || exit 1

