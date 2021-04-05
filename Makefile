panel:
	-mkdir -p "$$HOME"/www/moody
	(cd panel && make build && mv build/ "$$HOME"/www/moody)