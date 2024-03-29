# Project Name

YOYO-CLI

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [Testing](#testing)
- [Deployment](#deployment)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## Getting Started

Welcome! I'm excited to introduce you to "YOYO-CLI", an open-source Node.js tool that can help you boost productivity in your development projects. Say goodbye to frustrating setup issues - this CLI automates the entire process, from directory creation to npm setup. This means less manual errors and more time for other important things.
YOYO_CLI creates a Node.js project that supports TypeScript and ECMAScript with three environments: development, staging, and production.

By using "YOYO-CLI", we can ensure that all our Node.js projects follow good architecture. What's more, we can even add more architectures to the tool to generate projects that suit our specific needs.

But wait, there's more! In the future, "YOYO-CLI" will be able to handle feature generation seamlessly. Generating features will include controller layers, interfaces, enums, and service layers - all set up automatically.

And that's just the beginning! We have big plans ahead, including commands for generating CRUD APIs and more. This will make the development journey even smoother and more enjoyable.

Here's the deal: this project is open source, and we would love to have your contributions, suggestions, and ideas. Let's work together to make "YOYO-CLI" the best tool it can be!

### Prerequisites

Before you begin, ensure you have met the following requirements:

- [Node.js](https://nodejs.org/) and [Yarn](https://classic.yarnpkg.com/) installed on your system.
- Your Node.js version must support ECMAScript Modules (ESM).

### Installation

Step-by-step guide on how to set up and run the project locally.

## Project Structure

The project is organized as follows:

- `cmd/`: This folder contains files that provide the command-line functionality of the project. These files are responsible for executing various tasks and commands.

- `internal/`: This folder contains folders, each subfolder is a type of project architecture

- `internal/featureArch/`:This subfolder will generate a project based on featureArch

- `internal/featureArch/outterStrcture`:This subfolder will generate the outter strcuture of the project (files outside src folder)
- `internal/featureArch/innerStructure`:This subfolder will generate the inner structure of the project( files inside src folder)

## Usage

To create a new Node.js project with TypeScript support and multiple environments, follow these steps:

1. Open your terminal.

2. Run the following command to build the yoyo tool:

   ```bash
   chmod +x build.sh or ./build.sh
   ```


3. Go to your project location and type to create project
   ```bash
   yoyo-cli init
   ```

## CMD 

1. To create feature type 
   ```bash
   yoyo-cli g-feature
   ```
2. To create a provider
   ```bash
   yoyo-cli g-provider
   ```
## Contributing

Contributions are welcome! If you'd like to contribute to this project, follow these steps:

1. Fork the repository.

2. Create a new branch for your contribution:

   ```bash
   git checkout -b feature/your-feature-name
   ```

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

### Attribution

This project was started and developed by [Yesser Rebai](mailto:yesser.rebai96@gmail.com).

### Legal Notice

The use, reproduction, distribution, or modification of this software, either in source code or as a compiled binary, is subject to the terms and conditions of the MIT License. You can find a copy of the MIT License in the [LICENSE](./LICENSE) file.

By contributing to this project, you agree that your contributions will be licensed under the same terms and conditions as stated in the MIT License.

Please make sure to review and understand the license terms before using or contributing to this project.
