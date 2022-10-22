# Kaminoan

Easy to use cli tool for organize your repositories.

## About
* cli to wrap a normal `git clone <repository-url>` command and add some more useful logic

### cofig
* the config file is located at `<home-directory>/.kaminoan/settings.json`
* if this file doesn't exist you can simply use input prompt at first run
* otherwise create this file manually including f.e. this content `{"workspace":"/Users/<username>/workspace"}` 

### workspace
* if workspace is set all your repositories you clone this Kaminoan are stored under this location
* if workspace is empty stores the repository relative to your actual path

## Usage
* clone this repository
* for first tries: `go run main.go <repository-url>`

## Supported Url formats
* `https://<domain>/<groups>/<repo-name>.git`
* `git@<domain>:<groups>/<repo-name>.git`
