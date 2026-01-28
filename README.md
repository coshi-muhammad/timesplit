# timesplit
A time managment application for people who need a completly structured day or week 
## 🚀 Quick Start
### 1. Prerequisites

Because this project uses CGO, you must have a C compiler and graphics development headers installed on your system.

    Go: Version 1.19 or later.

    C Compiler: * Windows: MSYS2 (with mingw-w64-x86_64-toolchain).

        macOS: Xcode Command Line Tools (xcode-select --install).

        Linux: Graphics headers (e.g., libgl1-mesa-dev, xorg-dev on Ubuntu).
### 2. Installation
you can use one of the taged releases found in [releases]( "github.com/coshi-muhammd/timesplit/internal/core")

If you want to build from source run the following commands

Clone the repository and fetch the dependencies:

```bash
git clone https://github.com/coshi-muhammad/timesplit.git
cd yourproject
go mod tidy
```

To run the application immediately:
   ```bash
    go run ./cmd/timesplit
   ```

Or build an executable:
```bash
    go build ./cmd/timesplit
```


## 📝 License
Distributed under the MIT License.
