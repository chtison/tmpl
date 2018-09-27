.PHONY: version version-test

VERSION := 1.0.0

version:
	sed -Ei 's/^([[:space:]]+Version:[[:space:]]+)".*",$$/\1"$(VERSION)",/' main.go

version-test: version
	[ '$(VERSION)' = "`grep -E '^[[:space:]]+Version:' main.go | cut -d\\" -f2`" ]
