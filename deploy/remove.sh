docker rmi -f toan3082004/ltnc-be:latest
docker rmi -f ltnc-be:latest
docker rmi $(docker images -f "dangling=true" -q)