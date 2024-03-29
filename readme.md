# Pinjur Lunch 

The #1 world app for getting your lunch.

## Deploying:

#### Requirements:
- `go` installed and present in the `$PATH`
- `docker` installed
 
1. Run the `build.sh` script inside the `scripts` folder. It will create `batato-lunch` file inside the `cmd` directory.
2. `cd` into the `deployment` directory
3. run the `docker-compose up` command - This will start the app. Change the `docker-compose.yml` file to configure `port` and `host` address

### Stable version

Before building the project, make sure you are checking out the stable branch

#### Building the image (Dockerfile)

- from the project root:
```bash
docker build -t pinjur-lunch:stable -f build/Dockerfile-stable .
```

#### Deploying with docker-compose 

cd into deployment directory
```bash
cd deployment/
```
run the `docker-compose up` with different project name detach it(`-d`)
```bash
docker-compose -p pinjur-lunch up -d
```

## Goals:

- [ ] Create a pure docker based build script so `go` will no longer be a reqirement
- [ ] Implement a database (MongoDB) for storing order information as whatever
else might be necessary
- [ ] Reformat the JSON responses from the handlers to allow front end to be 
completely separated from the back end. 