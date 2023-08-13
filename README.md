# GohCMS

## 🚧 This project is under development !

The features are not complete, not fully tested and subject to many changes.

## Setup

Don't forget to setup the [Environment variables](#environment-variables)!

### Use with Docker

#### Compose

The easiest way to run the project is to use docker-compose.
You still need to set up the environment variables, but you can use the `.env` file for that.
Note that some supplementary variables are required for docker-compose to work, like `DOCKER_DB_FOLDER`.

```bash
docker compose up --build -d
```

### Build and run it yourself

#### Build

There are 2 scripts to build the project:

- `super_build.sh`
- `super_build.bat`

They will build the project for windows, linux and macOS. The binaries will be in the `bin` folder.
You can also really just use `go build` to build the project for your current OS.

#### Run

Once the project is built, you can simply run the binary for your OS. Setting the environment variables is
of course required, but not necessarily via the `.env` file.

### Environment variables

| Name                 | Type   | Description                             | Comment                                                                                               |
|----------------------|--------|-----------------------------------------|-------------------------------------------------------------------------------------------------------|
| ENVIRONMENT          | string | The environment the API is running in   | required, `development` or `production`                                                               |
| PORT                 | int    | The port the API will use               | required                                                                                              |
| CORS_ALLOWED_ORIGINS | string | The allowed origins for CORS            | required, semicolon separated list                                                                    |
| DB_FILE              | string | The path to the sqlite db file          | required, can be at the root but name still required (e.g. `./gohcms.db`) ; have to end up with `.db` |
| DOCKER_DB_FOLDER     | string | The path to the sqlite db file's folder | required with docker, this will basically be the host machine folder that contains the sqlite db file |

## API Usage

TODO

## Demo

TODO
