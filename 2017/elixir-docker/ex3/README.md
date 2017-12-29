git clone https://github.com/mremond/twelve.git

cd ~/elixir/twelve

docker build -t my_app .

docker run -it --rm -e NODE_NAME='node2' -e NODES='node1@MacBook-Pro-de-Mickael' my_app

docker-compose up
