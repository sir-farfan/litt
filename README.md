## Employee admin service

This is a local service for managing our employee database

The full technical requirements for this project so far can be seen in [REQS](PROJECT_REQS.md)

The overall DB structure is this

![DB!](./db/teams.png)

### Local development

In order to start the services needed you need to install Docker and Docker compose
(usually shipped together) in your development machine, there are a lot of resources
on how to do that.

Once the tool is up and running just run this in the project directory:
```bash
docker-compose up
```

It will start a local MySQL database and create the tables as defined in ./db/database.sql

### Building and publishing the service

Request access to this private repository to see the CI/CD pipelines  
https://github.com/sir-farfan/litt/actions

### Project dashboard:

https://github.com/sir-farfan/litt/projects/1

#### Restarting the DB

The DB running in Docker will keep all your records across restarts, if you want
to delete everything for whatever reason, just run

```bash
docker compose rm
```

## Running with docker

### Public repository

Published and nightly build images can be found in Docker Hub at:
https://hub.docker.com/repository/docker/sulfurf/litt