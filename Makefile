build:
	docker image build -f Dockerfile -t go-program .
run:
	docker container run --detach --name discord-bot -e TZ=Asia/Almaty go-program
