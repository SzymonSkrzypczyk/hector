# Hector
A static cv generator for developers who are tired of changing their only CV created in existance.

## Features
- Generate a professional CV in PDF format.
- Easy to customize with your personal information.
- Supports multiple predefined templates.
- Command-line interface for quick generation.
- Open-source and free to use.

## Installation
To install Hector, clone the repository and navigate to the project directory:
```bash
git clone git@github.com:SzymonSkrzypczyk/hector.git
cd hector
```
Build go project
```bash
go build
```

## Get help
To see the available commands and options, run:
```bash
./hector --help
```

## Structure
The project is structured as follows:
- 📁 [`cmd/`](cmd/): Contains the main application code and command-line interface.
- 📁 [`themes/`](themes/): Contains predefined CV templates.
- 📁 [`examples/`](examples/): Example CV data files.
- 📁 [`internal/parser`](internal/parser/): Package for parsing CV data files.
- 📁 [`internal/pdf`](internal/pdf/): Package for PDF generation utilities.
- 📁 [`internal/pipeline`](internal/pipeline/): Package for managing generation pipelines.
- 📁 [`internal/render`](internal/render/): Package for rendering CVs using templates.
- 📁 [`internal/schemas`](internal/schemas/): Package for defining CV data schemas.
- 📁 [`internal/theme`](internal/theme/): Package for managing CV themes.
- 📝 [`main.go`](main.go): The entry point of the application.