# Potato Lunch 

The #1 world app for getting your lunch.

## Deploying:

#### Requirements:
- `go` installed and present in the `$PATH`
- `docker` installed
 
1. Run the `build.sh` script inside the `scripts` folder. It will create `pinjur-lunch` file inside the `cmd` directory.
2. `cd` into the `deployment` directory
3. run the `docker-compose up` command - This will start the app. Change the `docker-compose.yml` file to configure `port` and `host` address
