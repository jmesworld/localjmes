# localjmes

One-click local JMES testnet and ecosystem for rapid prototyping

## About

LocalJMES allows to orchestrate with a single `docker-compose` file a complete JMES testnet.

## Prerequisites

- [Docker](https://www.docker.com/)
- [`docker-compose`](https://github.com/docker/compose)
- Supported known architecture: x86_64
- 16+ GB of RAM is recommended.

## Install LocalJMES

1. Run the following commands::

```sh
$ git clone --depth 1 https://www.github.com/jmesworld/localjmes
$ cd localjmes
```

2. Make sure your Docker daemon is running in the background and [`docker-compose`](https://github.com/docker/compose) is installed.

## Start

```sh
$ docker-compose up
```

Your environment now contains:

- [jmesd](http://github.com/jmesworld/core) RPC node running on `tcp://:26657`  
  - It will generate a local network `jmes-888`. As a testnet, it will keep a faucet funded, setup a validator called dev-masternode.  
- LCD (light client daemon provides a REST-based adapter for the RPC endpoints) running on http://:1317  
- [FCD](http://www.github.com/terra-money/fcd) (full client daemon) running on http://localhost:3060  
- An instance of the backend server managing the identity, items and files for the application.  
- An instance of the blockexplorer.  
- Proxy from 26657 to 1889  
- Faucet running on :3002 (use jmes.js to easily request credit or call :3002/credit?address={ADDRESS})  

## Stop

```sh 
$ docker-compose stop
```

## Reset

```sh
$ docker-compose rm
```
