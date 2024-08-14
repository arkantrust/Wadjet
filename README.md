# wadjet

Wadjet is a lightweight, open-source system monitoring tool designed to keep an eye on your server's critical metrics including CPU, memory, network speed, disk usage, and more. Built with efficiency in mind, Wadjet ensures minimal resource consumption while providing valuable insights into your system's performance. Future plans include transitioning to the QUIC protocol for enhanced speed and security.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Features](#features)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [Roadmap](#roadmap)
- [License](#license)

## Installation

To install Wadjet, you can either download the pre-built binary or build it from source, keep in mind you can do either in both your servers and your local machine.

### From Source

To build Wadjet from source, you will need to have Go installed on your machine. If you don't have it installed, you can download it from the [official website](https://go.dev/doc/install).

```bash
git clone https://github.com/arkantrust/wadjet.git
cd wadjet
go build -o wadjet
./wadjet --help
```

### Pre-built Binary

TODO

## Usage

To start Wadjet, simply run the binary and pass the desired flags. For example, to view server RTT, you can run the following command:

```bash
./wadjet -h server.example.com -p 6000 rtt
# Output: RTT: 0.5ms
```

## Commands

Wadjet supports the following commands:

- `interfaces`: Lists all network interfaces on the server.
- `ip`: Returns the server's IP address.
- `time`: Returns the server's local time.
- `rtt`: Measures round-trip time for a 1024-byte message in ms.
- `speed`: Measures transmission speed for an 8192-byte message in kB/s.

## Features

- **Monitor network interfaces and IP configuration***

- **Measure RTT and transmission speed***

- **Lightweight and minimal resource usage***

- **Cross-platform Support***

> *: active development

See the [Roadmap](#roadmap) for more information on planned features.

## Configuration

Wadjet can be configured using command-line flags or a configuration file. The configuration file should be named `wadjet.yaml` and placed in the same directory as the binary. Here is an example configuration file:

```yaml
server:
  host: server.example.com
  port: 6000
```

You can then run Wadjet with the `-c` flag to specify the configuration file:

```bash
./wadjet -c wadjet.yaml rtt
```


## Contributing

Contributions are welcome! Please check out our [contributing guidelines](CONTRIBUTING.md) before submitting a pull request.

#### Roadmap

- [ ] **Full system monitoring**: Monitor CPU, memory, disk usage, and more.

- [ ] **Customizable Alerts**: Set custom alerts for your server's critical metrics to receive notifications when they exceed a certain threshold.

- [ ] **Real-time Monitoring**: Monitor your server's performance in real-time with Wadjet's live updates.

- [ ] **QUIC Protocol Support**: Transition to the QUIC protocol for enhanced speed and security.

- [ ] **Web Interface**: Access Wadjet's monitoring features through a web interface.

- [ ] **Docker Support**: Run Wadjet in a Docker container for easy deployment.

- [ ] **API Integration**: Integrate Wadjet with your existing monitoring tools through an API.

- [ ] **Custom Plugins**: Create custom plugins to extend Wadjet's functionality.

- [ ] **User Authentication**: Secure your monitoring data with user authentication.

- [ ] **Data Visualization**: Visualize your server's performance metrics with graphs and charts.

- [ ] **Historical Data**: View historical data for your server's performance metrics.

- [ ] **Custom Dashboards**: Create custom dashboards to monitor your server's performance.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

Built with ❤️ using [Go](https://go.dev/).

Inspired by [Prometheus](https://prometheus.io/) and [Netdata](https://www.netdata.cloud/).