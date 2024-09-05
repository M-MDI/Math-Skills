#!/usr/bin/env bash

IMAGE="stat-bin"
VALID_BIN=$(ls -xm ./bin | sed 's:, :|:')

help() {
	cat <<EOF
Run porgram as follow 
./run.sh [$VALID_BIN]
EOF
}

exec_stat_bin() {
	if echo "$VALID_BIN" | grep -q "$1"; then
		docker run -v "$(pwd):$(pwd)" -w "$(pwd)" --rm "$IMAGE" "/stat-bin/$1"
	else
		echo "Error: binary not found!"
		help
		exit 1
	fi
}

if [[ $# -ne 1 ]]; then
	help
	exit 1
fi

if ! docker images | grep -q "$IMAGE"; then
	docker build -t "$IMAGE" .
fi

exec_stat_bin "$1"
