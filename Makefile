build:
	cd react && yarn build
	cd react && mv build ../
	go build
	mv -b goeubyt build