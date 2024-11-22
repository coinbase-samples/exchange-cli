# Exchange CLI README

## Overview

The Exchange CLI is a sample Command Line Interface (CLI) application that generates requests to and receives responses from [Coinbase Exchange's](https://exchange.coinbase.com/) [REST APIs](https://docs.cdp.coinbase.com/exchange/docs/welcome). The Exchange CLI is written in Go, using [Cobra](https://github.com/spf13/cobra).

## License

The Exchange CLI is free and open source and released under the [Apache License, Version 2.0](LICENSE.txt).

The application and code are only available for demonstration purposes.

## Usage

To begin, navigate to your preferred directory for development and clone the Exchange CLI repository and enter the directory using the following commands:

```bash
git clone https://github.com/coinbase-samples/exchange-cli
cd exchange-cli
```

Next, pass an environment variable via your terminal called `EXCHANGE_CREDENTIALS` with your API information.

Exchange API credentials can be created in the Exchange web console under Settings -> APIs.

`EXCHANGE_CREDENTIALS` should match the following format:

```bash
export EXCHANGE_CREDENTIALS='{
"apiKey":"api_key_here",
"passphrase":"passphrase_here",
"signingKey":"signing_key_here",
}'
```

You may also pass an environment variable called `exchangeCliTimeout` which will override the default request timeout of 7 seconds. This value should be an integer in seconds.

Build the application binary and specify an output name, e.g. `exctl`:

```bash
go build -o exctl
```

To ensure your project's dependencies are up-to-date, run:

```bash
go mod tidy
```

To make your application easily accessible from any location, move the binary you created to a directory that's already in your system's PATH. For example, these are the commands to move `exchangectl` to `/usr/local/bin`, as well as set permissions to reduce risk:

```bash
sudo mv exctl /usr/local/bin/
chmod 755 /usr/local/bin/exctl
```

To verify that your application is installed correctly and accessible from any location, run the following command. It will include all available requests:

```bash
exctl
```

Finally, to run commands for each endpoint, use the following format to test each endpoint. Please note that many endpoints require flags, which are detailed with the `--help` flag.

```bash
exctl list-accounts
```

```bash
exctl create-order --help
```

```bash
exctl get-product-book -p ETH-USD -l 1
```

