build:
	docker image build -f Dockerfile -t go-program .
run:
	docker container run --detach --name discord-bot go-program
