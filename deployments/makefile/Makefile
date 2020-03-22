setup:
	if [ -e ./dist ]; then rm -rf ./dist; fi; mkdir ./dist; mkdir -p ./dist/darwin/bin; mkdir -p ./dist/linux/bin

build:
	@echo "******************"
	@echo "**** Building ****"
	@echo "******************"

	GO111MODULE=on GOOS=linux go build -mod=vendor ../../cmd/main.go;cp dist/linux/bin/main ../docker;mv ./main dist/linux/bin/main


docker-build-prep:
	cd docker;cp ../dist/linux/bin/main .

copy-migration:
	cd ../docker; cp ../../pkg/migartion_scripts/* .

docker-build-cleanup:
	cd ../docker;rm main

delete-migration:
	cd ../docker; rm -rf *init* .

release-tag: build

	cd ../; git tag `cat .version`
	git push --tags

release: release-tag copy-migration
	cd ../docker; docker build -t "singaravelan21/todo_list_user_srv" . --no-cache;
	cd ../docker; docker tag "singaravelan21/todo_list_user_srv" "singaravelan21/todo_list_user_srv:`cat ../.version`";
	cd ../docker; docker push "singaravelan21/todo_list_user_srv:`cat ../.version`"
	make docker-build-cleanup
	make delete-migration

