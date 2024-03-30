# RPCRequest

RPCRequest is a versatile command-line utility designed for executing JSON-RPC (Remote Procedure Call) requests. 
While it's tailored for seamless interaction with Ethereum RPC endpoints, its flexible design makes it suitable for a wide range of JSON-RPC based services.

## Features

- **Ethereum RPC Support:** Optimized for Ethereum, facilitating direct command-line interactions with Ethereum nodes.
- **Broad JSON-RPC Compatibility:** Use it with any JSON-RPC service, not just Ethereum Nodes.
- **Environment Variable & Command Line URL Configuration:** Specify your RPC node URL through an environment variable or directly as a command-line argument.
- **Basic Auth Support:** Easily pass basic authentication credentials within the URL for secure endpoints.
- **Flexible Parameter Typing:** Convert string arguments to booleans, integers, or hex-encoded strings to meet various parameter requirements.

## Getting Started

### Installation


#### Using Homebrew

```bash
brew install draganm/tools/rpcrequest
```

#### Using `go install`

```bash
go install github.com/draganm/rpcrequest
```

### Usage

RPCRequest accepts the JSON-RPC method name and parameters as command-line arguments. 
The first argument should always be the method name, followed by each parameter as a separate argument.

#### Example Command

```bash
rpcrequest --node-url https://my-node/path eth_getBlockByNumber asHex:12345 bool:true
```

### Specifying the RPC Node

You can specify the RPC node URL in two ways:

1. **Environment Variable:** Set the `NODE_URL` environment variable with your RPC node URL.
2. **Command-Line Flag:** Use the `--node-url=<url>` flag to provide the URL directly in your command.

### Authentication

If the RPC endpoint requires basic authentication, include the credentials directly in the URL as follows:

```url
https://username:password@hostname/path
```

### Converting to JSON Types

RPCRequest can automatically convert string arguments into the required JSON types. Use the following prefixes to indicate the type:

- `bool:` Converts the following value to a boolean (`true` or `false`).
- `int:` Converts the following value to an integer.
- `asHex:` Converts an integer value into a hex-encoded string prefixed with `0x`. This is particularly useful for Ethereum RPC parameters like block numbers.

## Contributing

Contributions are welcome, send your issues and PRs to this repo.

## License

[MIT](LICENCE) - Copyright Dragan Milic and contributors.