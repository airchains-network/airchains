# Airchains Modular ZK Framework
A versatile and powerful framework for creating customized rollups with a variety of options. Our framework supports EVM, SVM, and CosmWasm-based rollups, with data availability choices including Celestia and Avail. Secured with cutting-edge zk-proofs, the framework integrates seamlessly with Airchains' settlement chain, offering a robust solution for modular blockchain development needs.

## Features
- **Support for Multiple Virtual Machines**: Create rollups based on Ethereum Virtual Machine (EVM), Solana Virtual Machine (SVM), or CosmWasm.
- **Data Availability Options:** Choose from leading data availability layers like Celestia and Avail.
- **Security with zk-Proofs:** Enhanced security and efficiency through zero-knowledge proofs.
- **Integration with Airchains:** Seamless settlement on Airchains' settlement chain.

## Prerequisites
Before you begin, ensure you have the following installed:
- [Golang](https://golang.org/doc/install) (version 1.20+)
- [Rustc](https://www.rust-lang.org/tools/install) (version 1.74+)
- [Cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html) (version 1.74+)
- [jq](https://stedolan.github.io/jq/download/)
- [Node.js](https://nodejs.org/en/download/) (version 18+)

## Installation & Run
Clone the repository and install the dependencies:

```bash
git clone https://github.com/airchains-network/airchains.git
cd airchains
```
Build the project:
```bash
go build -o airchains
```
Run the project:
```bash
./airchains start
```

## License
This project is licensed under the [LICENSE](./LICENSE) - see the file for details.
