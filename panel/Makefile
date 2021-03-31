include .env

node_modules : package.json package-lock.json
	npm install

dev : node_modules
	npm start

build : node_modules
	npm run build
clean :
	rm -rf node_modules
	rm -rf build

.PHONY : dev
.PHONY : start
.PHONY : stop
.PHONY : clean
