# *moody*

An open-source smart home hub by antima.it

## Contents

- [Installation](#installation)
    - [Run via docker-compose](#run-via-docker-compose)
    - [Run via make](#run-via-make)
## Installation

Clone the repo and cd into it:

```bash
git clone https://github.com/Abathargh/moody-go
cd moody-go
```

### Run via docker-compose:

Once you have initialized the configuration files and the certificates, you can start using moody through docker,
interfacing with the admin panel reachable from http://localhost:3000.

```bash
docker-compose up --build -d
```


Pre-built images for each service are available at https://hub.docker.com/u/abathargh.
More instructions about every feature can be found in each subfolder.


### Run via make

If your backend services are hosted on another machine, where the api gateway is reachable, you can build and run
the front-side (broker + gateway + webapp) via make. In this case you will need to install nodejs, npm mosquitto and
golang, then run:

```bash
make build-front
make run-front

# Stop with
make stop-front
```