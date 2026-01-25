# Hector

![CI](https://github.com/SzymonSkrzypczyk/hector/actions/workflows/ci-checks.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/SzymonSkrzypczyk/hector)](https://goreportcard.com/report/github.com/SzymonSkrzypczyk/hector)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/release/SzymonSkrzypczyk/hector.svg)](https://github.com/SzymonSkrzypczyk/hector/releases)

A static cv generator for developers who are tired of changing their only CV created in existance.

![logo](docs/assets/logo_transparent.png)

## Features
- Generate a professional CV in PDF format.
- Easy to customize with your personal information.
- Supports multiple predefined templates.
- Command-line interface for quick generation.
- Open-source and free to use.

## Installation

To install Hector, clone the repository and build the project:

```bash
git clone git@github.com:SzymonSkrzypczyk/hector.git
cd hector
go build
```

Alternatively you can use `task` to build the project.
First install task from [here](https://taskfile.dev/installation/).
Then run:
```bash
task build
```

For more detailed installation instructions, including troubleshooting, please refer to the [User Guide](docs/USER_GUIDE.md).

## Usage

### Quick Start
Once built, you can generate a CV immediately using the provided examples.

```bash
# Generating a PDF using the 'normal' theme
./hector generate --theme normal --data examples/normal_example.yaml --output-dir .
```

To see the available commands and options, run:
```bash
./hector --help
```

For a comprehensive guide on all commands, check the [User Guide](docs/USER_GUIDE.md).

## Configuration
Hector uses a YAML file to store your resume data. You can find example configurations in the [`examples/`](examples/) directory.

Key sections in the configuration:
- `candidate_info`: Personal details and summary.
- `contact`: Email, phone, and links.
- `experience`: Work history.
- `education`: Academic background.
- `skills`: Technical keywords.

To validate your configuration file:
```bash
./hector validate --data your_cv_data.yaml
```

## Structure
The project is structured as follows:
- 📁 [`cmd/`](cmd/): Contains the main application code and command-line interface.
- 📁 [`themes/`](themes/): Contains predefined CV templates.
- 📁 [`examples/`](examples/): Example CV data files.
- 📁 [`internal/`](internal/): Core application logic (parser, pdf, pipeline, etc.).
- 📝 [`main.go`](main.go): The entry point of the application.

## Documents
- [`USER_GUIDE.md`](docs/USER_GUIDE.md): Detailed usage and installation guide.
- [`CONTRIBUTING.md`](docs/CONTRIBUTING.md): Guide for contributors.
- [`CHANGELOG.md`](docs/CHANGELOG.md): History of changes.
- [`CODE_OF_CONDUCT.md`](docs/CODE_OF_CONDUCT.md): Code of conduct.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
