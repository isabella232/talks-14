mix phx.server

docker build -t parisex2 .

docker run -it --rm -p 4000:4000 parisex2

docker-compose up

docker-compose up --scale myapp=3

docker-compose rm