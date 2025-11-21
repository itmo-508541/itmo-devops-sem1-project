#!/bin/bash

$SERVER_ADDRESS=$1
while ! nc -z -v -w30 $SERVER_ADDRESS 22; do
  sleep 1
done

ssh runner@$SERVER_ADDRESS -o StrictHostKeyChecking=no "curl -fsSL https://get.docker.com -o get-docker.sh ;
  sudo bash ./get-docker.sh ;
  sudo chown runner:runner /srv ;
  rm /srv/deploy.yml || true"

mkdir --parent ./build
docker compose -f docker-compose.deploy.yml config > ./build/deploy.yml
scp -P 22 -r ./build/deploy.yml runner@$SERVER_ADDRESS:/srv/deploy.yml

ssh runner@$SERVER_ADDRESS -o StrictHostKeyChecking=no "cd /srv ;
  docker compose stop || true ;
  rm docker-compose.yml || true ;
  mv deploy.yml docker-compose.yml ;
  docker compose pull ;
  docker compose up -d"
