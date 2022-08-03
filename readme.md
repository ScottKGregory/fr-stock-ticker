# Forge Rock Stock Ticker

This project creates an API surface that calls through to the Alphavantage API to retrieve daily timeseries data according to the following config:

- `FR_DAYS`: The number of days worth of data to return
- `FR_SYMBOL`: The stock symbol to lookup data for
- `FR_API_KEY`: The API key that will be used as part of requests to the Alphavantage API

_All config can be set via environment variables, and are currently all required_

## Building the docker image

The root of the project contains a dockerfile to allow for building an image for the stock ticker app. The image will be built based on the Google's [distroless](https://github.com/GoogleContainerTools/distroless) image and will therefore contain only the absolute essentials to run the app.

To build the image navigate to the root of the project and run `docker build --platform=linux/amd64 -t scottgregory/fr-stock-ticker:v0.0.1 .`.
It can then be run locally using docker via `docker run -p 3000:3000 -e FR_DAYS=5 -e FR_SYMBOL=MSFT -e FR_API_KEY=<SECRET> fr-stock-ticker:v0.0.1`.
To push the image to it's repo on [docker hub](https://hub.docker.com/repository/docker/scottgregory/fr-stock-ticker/general) run `docker push scottgregory/fr-stock-ticker:v0.0.1`

## Deploying the application

There are deployment manifests in the `/deploy` folder, these can be modified and applied as appropriate. To allow for better configuration and reuse something like Helm of Kustomize may be worth using rather than vannila manifest files.

To deploy:

1. Update the API key in `deploy/secret.yaml` to the correct value, it will need to be base64 encoded
2. Run `kubectl apply -f deploy`
