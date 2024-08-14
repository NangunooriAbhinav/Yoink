# Yoink

**Yoink** is a simple command-line tool for converting images between different formats. It currently supports converting images to JPEG and PNG formats.

## Features

- Convert images to JPEG or PNG format.
- Easy-to-use command-line interface.

## Installation

To install **Yoink**, follow these steps:

### Prerequisites

- **Go** must be installed on your system. You can download and install it from [the official Go website](https://golang.org/dl/).

### Installation Steps

1. **Clone the Repository (Optional)**

   If your project is hosted in a Git repository, clone it to your local machine:

   ```bash
   git clone https://github.com/yourusername/yoink.git
   cd yoink
   ```

2. **Build and Install**

   Run the provided installation script to compile the Go code and install the `yoink` command globally on your system:

   ```bash
   ./install.sh
   ```

   This script will:

   - Compile the Go program into an executable named `yoink`.
   - Move the executable to `/usr/local/bin` to make it accessible from anywhere in your terminal.
   - Ensure the executable has the correct permissions.

### Manual Installation (Alternative)

If you prefer to manually install the tool without using the installation script:

1. **Build the Go Program**

   ```bash
   go build -o yoink
   ```

2. **Move the Executable to a Directory in Your `PATH`**

   ```bash
   sudo mv yoink /usr/local/bin/
   ```

3. **Ensure Executable Permissions**

   ```bash
   sudo chmod +x /usr/local/bin/yoink
   ```

## Usage

After installation, you can use the `yoink` command to convert images.

### Basic Command Syntax

```bash
yoink [file path] [-j/-p]
```

### Options

- `-j`: Convert the image to JPEG format.
- `-p`: Convert the image to PNG format.

### Examples

1. **Convert an Image to JPEG**

   ```bash
   yoink /path/to/image.png -j
   ```

   This will convert `image.png` to `image.jpeg`.

2. **Convert an Image to PNG**

   ```bash
   yoink /path/to/image.jpg -p
   ```

   This will convert `image.jpg` to `image.png`.

### Help

To display help information:

```bash
yoink -h
```

## Contributing

If you'd like to contribute to **Yoink**, please feel free to submit a pull request or open an issue on the project's GitHub page.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

---

This README file provides a clear explanation of the `yoink` tool, including how to install it, how to use it, and where to get help. You can save this content into a `README.md` file in your project's root directory.
