# Hector User Guide

**Hector** is a static CV generator that helps you turn a YAML configuration file into a professional PDF resume using various themes.

---

## 🛠️ Installation

### Prerequisites

To run Hector, you need **Go (Golang)** installed on your system.

#### 1. Install Go
*   **Windows / Mac**: Download the installer from [go.dev/dl](https://go.dev/dl/) and run it.
*   **Linux / WSL**:
    ```bash
    sudo apt update
    sudo apt install golang-go
    ```
    *Verify installation:* `go version`

#### 2. Clone & Build
Download the Hector source code and build the application.

```bash
# Clone the repository
git clone https://github.com/SzymonSkrzypczyk/hector.git
cd hector

# Download dependencies
go mod download

# Build the executable
# Windows:
go build -o hector.exe

# Linux / Mac / WSL:
go build -o hector
```

---

## 🚀 Quick Start

Once built, you can generate a CV immediately using the provided examples.

```bash
# Generating a PDF using the 'normal' theme
./hector generate --theme normal --data examples/normal_example.yaml --output-dir .
```

*Note: On Linux/WSL, replace `./hector` with `./hector` or just `hector` if added to PATH. On Windows powerhshell use `.\hector.exe`*

---

## 📖 Command Line Reference

Hector provides several commands to help you manage your CV generation.

### 1. `generate`
The core command to build your PDF resume.

**Usage:**
```bash
./hector generate --data <file> [flags]
```

**Flags:**
*   `-d, --data` *(required)*: Path to your CV content file (YAML or JSON).
*   `-t, --theme`: Theme to use (default: `normal`).
*   `-o, --output-dir`: Directory where the generated PDF will be saved.

### 2. `preview`
View your CV in a web browser before generating the final PDF.

**Usage:**
```bash
./hector preview --data <file> [flags]
```

**Flags:**
*   `-d, --data` *(required)*: Path to your CV content file.
*   `-t, --theme`: Theme to use.
*   `--watch`: Watch for file changes and auto-reload the preview.
*   `--no-open`: Start the server but don't open the browser automatically.

### 3. `validate`
Check your YAML data file for errors or missing fields required by a theme.

**Usage:**
```bash
./hector validate --data <file> [flags]
```

**Flags:**
*   `-d, --data` *(required)*: Path to your CV content file.
*   `-t, --theme`: The theme schema to validate against (default: `normal`).

### 4. `themes`
List all valid themes currently available in the `themes/` directory.

**Usage:**
```bash
./hector themes
```

### 5. `template`
Extract the raw YAML template structure for a specific theme. This is useful if you want to see exactly what fields a theme supports.

**Usage:**
```bash
./hector template --theme <name> --file <output_path>
```

---

## 📝 Creating Your CV (YAML Format)

Hector uses a YAML file to store your resume data. Below is a breakdown of the standard structures.

### Basic Structure

```yaml
path: "photo.jpg"  # Path to your profile picture

candidate_info:
  name: "John Doe"
  about: "A brief professional summary..."

contact:
  email: "john@example.com"
  phone: "+1 234 567 890"

skills:
  - Go (Expert)
  - Docker (Proficient)

languages:
  - language: English
    level: C2

# ... see examples/normal_example.yaml for full list
```

### Sections

*   **`candidate_info`**: Your name and "About Me" blurb.
*   **`contact`**: Email, phone, and other contact details.
*   **`experience`**: Work history (Role, Company, Period, Responsibilities).
*   **`education`**: Academic background (Institution, Degree, Period, Details).
*   **`projects`**: Side projects or major achievements.
*   **`certifications`**: List of earned certificates.
*   **`extra_activities`**: Volunteering, organizations, or hobbies.

---

## ❓ Troubleshooting

**"go: command not found"**
*   Ensure Go is installed (`go version`).
*   On WSL, check your `$PATH`. You might need to add `/usr/local/go/bin`.

**"Validation failed"**
*   Run `hector validate` to see exactly which fields are missing or malformed in your YAML file.
