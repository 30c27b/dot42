#!/bin/sh

# utils
join_args() {
	local arg
	printf "%s" "$1"
	shift
	for arg in "$@"; do
		printf " "
		printf "%s" "${arg// /\ }"
	done
}

abort() {
	printf "Error: %s\n" "$1"
	exit 1
}

execute() {
	if ! "$@"; then
		abort "$(printf "could not execute command: %s" "$(join_args "$@")")"
	fi
}

# check if runing os macOS
if [ "$(uname)" != "Darwin" ]; then
	abort "this script only works on macOS"
fi

# launch dot42 load command
execute "$HOME/.42/bin/dot42" "load"

# sleep for 10 seconds to prevent launchd from thinking the daemon has crashed
execute "sleep" "10"

# done
execute "echo" "dot42 has loaded your configuration."
