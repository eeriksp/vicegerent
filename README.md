# Vicegerent

<a href="https://snapcraft.io/vicegerent">
  <img alt="vicegerent" src="https://snapcraft.io/vicegerent/badge.svg" />
</a>

<img src="https://raw.githubusercontent.com/eeriksp/vicegerent/master/brand/logo.png" alt="Vicegerent logo" height="400px">

Vicegerent is a small Go application which helps you to run predefined sets of commands in the server invoked by an HTTP call.

The original intent behind the project was to trigger the update commands for Docker containers from a CI/CD script without giving SSH access to the server.

## Setup

### Preparing the configuration file

The configuration file should be located in `/etc/vicegerent/config.yml`. An example configuration could look like this:

```yaml
address: ':6060'
tasks:
  - name: 'Show software versions'
    url: 'versions'
    workDir: '/home/username'
    commands:
      - [ 'go', 'version' ]
      - [ 'python', '--version' ]
  - name: 'Update the application'
    url: 'update-application'
    workDir: '/home/username/docker-application'
    commands:
      - [ 'docker-compose', 'down' ]
      - [ 'docker-compose', 'pull' ]
      - [ 'docker-compose', 'up', '-d']
```

### Installing Vicegerent

#### Option 1: Install as a snap package

<a href="https://snapcraft.io/vicegerent">
  <img alt="Get it from the Snap Store" src="https://snapcraft.io/static/images/badges/en/snap-store-black.svg" />
</a>

1. Install the snap

   `sudo snap install vicegerent --classic`

   The demon will start automatically, although I won't see any output.

2. If you want to double-check that the demon started, use

   `sudo snap services vicegerent`

   You should see a line with service `vicegerent.daemon`

3. If you make any changes to the configuration file, restart the daemon with

   `sudo snap restart vicegerent`

#### Option 2: regular Go project

1. Download the source:
   
   `git clone git@github.com:eeriksp/vicegerent.git`

2. Run the application:

   ```
   cd vicegerent/src
   go run .
   ```
   You should see the following output:
   ```
   Ready to accept tasks. Listening on :6060 for
   - Show software versions: /task/versions
   - Update the application: /task/update-aplication
   ```

3. To run the process in the background, use

   `nohup go run . &`



### Check that the application is running properly

Go to `localhost:6060/task/versions` (adjust the port and URL to mach your configuration file). You should see something like this:

```
Started task "Show software versions"


--- Start to execute [go version] ---
go version go1.16.2 linux/amd64

--- Start to execute [python --version] ---
Python 3.8.5

Task finished.
```

## Platform support

Tested on Ubuntu 20.

If you can confirm that it is (not) working on some other platform, please let us know by opening an issue or submitting a pull request.

## Contributing

Contribution is welcome. Please document your changes in `CHANGELOG.md` when submitting a pull request.
