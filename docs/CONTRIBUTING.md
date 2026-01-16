# Contributing to Hector

Thank you for your interest in improving **Hector**! 🚀
Whether you're fixing a bug, adding a new CV theme, or improving the documentation, I welcome your contributions.

---

## ✨ Getting Started

### Prerequisites

You need to have [Go](https://golang.org/doc/install) installed on your machine.
I recommend using the latest stable version of Go.

### 1. Fork and Clone

1. Fork the repository on GitHub.
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/hector.git
   cd hector
   ```
3. Add the upstream repository to keep your fork in sync:
   ```bash
   git remote add upstream https://github.com/SzymonSkrzypczyk/hector.git
   ```

### 2. Install Dependencies

Hector uses Go modules. Run the following to download dependencies:

```bash
go mod download
```

---

## 🛠 Development Workflow

### Project Structure

- **`cmd/`**: Contains the main CLI application entry points.
- **`internal/`**: Core logic packages:
  - `parser/`: Parsing CV data files.
  - `pdf/`: PDF generation utilities.
  - `render/`: Template rendering logic.
  - `theme/`: Theme management.
  - `schemas/`: Data definitions.
- **`themes/`**: Predefined CV templates. New themes go here.
- **`examples/`**: Example CV data files (YAML/JSON).

### Running Locally

To build and run the project locally:

```bash
# Build the binary
go build

# Run the help command
./hector --help
```

### Running Tests

Please ensure various tests pass before submitting a PR.
Since `GOROOT` might vary, ensure your generic Go environment is set up successfully.

```bash
# Run all unit tests
go test ./...
```

### Code Style

Hector follows standard Go idioms. Please run `go fmt` before committing changes:

```bash
go fmt ./...
```

---

## 🎨 Adding a New Theme

If you want to contribute a new CV design:

1. Create a new directory in `themes/` with your theme name.
2. Add your HTML templates and CSS files.
3. Create a new theme struct that implements `ThemeSchema` interface.
4. Make sure to add your theme to `internal/theme/select_theme.go` to make it selectable.
5. Add new unit tests to `internal/theme/theme_test.go`.
6. Test it by generating a CV using your new theme:
   ```bash
   ./hector generate --theme=your_new_theme --data=examples/cv.yaml
   ```
7. Check if the generation works for multiple pages.

---

## 🚀 Submitting a Pull Request

1. Create a new branch for your feature or fix:
   ```bash
   git checkout -b feature/my-new-feature
   ```
2. Commit your changes with clear, descriptive messages.
3. Push to your fork and open a Pull Request against the `main` branch of the upstream repository.
4. Describe your changes detailedly in the PR description.

---

## 💬 Need Help?

If you have questions, feel free to open an issue or start a discussion in the [GitHub repository](https://github.com/SzymonSkrzypczyk/hector).

Happy coding! 🎉
