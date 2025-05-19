CONTAINER=alf-htmx
IMAGE=alloydflanagan-htmx:latest

build:
	docker build . --tag $(IMAGE)

shell:
	docker exec -it $(CONTAINER) /bin/bash

run:
	docker run -it --name $(CONTAINER) --rm -p 8000:80 $(IMAGE)

stop:
	docker stop $(CONTAINER)
