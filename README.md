# Protoc Plugin Generator for Redis RPC

[![Protoc Gen RPC Redis](https://github.com/ksysoev/protoc-gen-rpc-redis/actions/workflows/main.yml/badge.svg)](https://github.com/ksysoev/protoc-gen-rpc-redis/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/ksysoev/protoc-gen-rpc-redis/graph/badge.svg?token=WT9WBNMMZ3)](https://codecov.io/gh/ksysoev/protoc-gen-rpc-redis)

This package provides a plugin for the `protoc` utility, enabling the generation of Go code for [Redis RPC](http://github.com/ksysoev/redis-rpc) from Protocol Buffers (`.proto`) files. The plugin is designed to streamline the process of creating Go structs and service interfaces based on your Protocol Buffers definitions.


## Installation

```sh
go install github.com/ksysoev/protoc-gen-rpc-redis@latest
```

## Usage

To generate Go code from your `.proto` files using the custom `rpc-redis` plugin, run the following command:

```sh
protoc --rpc-redis_out=. --rpc-redis_opt=paths=source_relative echosrv.proto
```

An example of the generated code can be found in [this repository](https://github.com/ksysoev/echosrv). This example demonstrates how the plugin converts `.proto` definitions into Go code, including structs and service interfaces.

Feel free to explore the repository to see the plugin in action and understand how to integrate it into your own projects.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you encounter any problems or have suggestions for improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
