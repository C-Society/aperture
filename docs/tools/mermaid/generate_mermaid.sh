#!/bin/bash

gitroot=$(git rev-parse --show-toplevel)
export gitroot

docsdir=$gitroot/docs
export docsdir

GREP="grep"
SED="sed"
AWK="awk"
FIND="find"
# check whether we are using macOS
if [ "$(uname)" == "Darwin" ]; then
	GREP="ggrep"
	SED="gsed"
	AWK="gawk"
	FIND="gfind"
fi
export GREP
export SED
export AWK
export FIND

# accept a --force flag to force generation of all mermaid files
if [[ -n "${1:-}" ]] && [[ "$1" == "--force" ]]; then
	force=true
else
	force=false
fi
export force

# find all mmd files and generate svg and png files only when mmd contents change (using md5sum)
function generate_mermaid_images() {
	set -euo pipefail
	mmd_file="$1"
	# generate svg and png files only when mmd contents change (using md5sum)
	#shellcheck disable=SC2016
	md5sum=$(md5sum "$mmd_file" | $AWK '{print $1}')
	#shellcheck disable=SC2016
	md5sum_file=$(cat "$mmd_file".md5sum 2>/dev/null)
	if [ "$md5sum" != "$md5sum_file" ] || [ "$force" == true ]; then
		echo "generating svg and png files for $mmd_file"
		# generate svg and png files
		# loop formats svg and png
		# shellcheck disable=SC2043
		for fmt in svg; do #png; do
			npx -p @mermaid-js/mermaid-cli mmdc \
				--quiet --input "$mmd_file" --configFile "$docsdir"/tools/mermaid/mermaid-theme.json --cssFile "$docsdir"/tools/mermaid/mermaid.css --scale 2 --output "$mmd_file"."$fmt" --backgroundColor transparent
		done
		# update md5sum
		echo "$md5sum" >"$mmd_file".md5sum
	fi
}

export -f generate_mermaid_images

parallel -j4 --no-notice --bar --eta generate_mermaid_images ::: "$($FIND "$docsdir"/content -type f -name "*.mmd")"
