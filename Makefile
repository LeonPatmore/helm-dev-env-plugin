uninstall:
	helm plugin uninstall cool || true

install: uninstall
	echo "Ensure you are running as admin!"
	helm plugin install .

test:
	helm cool
