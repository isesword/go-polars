# go-polars

<p align="center">
    <img src="docs/assets/images/go-rust.png" width="300"/>
</p>

This project creates Go bindings for the Polars data manipulation library!

- 中文文档: [README_zh.md](README_zh.md)

## 🐻‍❄️ What is Polars?
Polars is an open-source library for data manipulation, known for being one of the fastest data processing solutions on a single machine. It features a well-structured, typed API that is both expressive and easy to use.

https://github.com/pola-rs/polars

## 📦 Installation

> [!NOTE]
> **Build Process & Security Considerations**
>
> The GitHub Actions runners cannot compile the Polars Rust bindings due to resource constraints, so binaries are currently compiled on a local development machine. While this isn't ideal from a security perspective, we've implemented several measures to ensure transparency and verifiability:
>
> - **🔍 Reproducible builds**: All build scripts are available in [`./scripts`](./scripts) for review
> - **🔐 Checksum verification**: Each binary release includes SHA256 and MD5 checksums
> - **📋 Build transparency**: Release notes include build environment details and dependency versions
> - **🏗️ Self-compilation option**: You can always build from source using `./build.sh`
>
> **To verify a binary download:**
> ```bash
> # Download the checksum file and verify
> sha256sum -c libpolars_go-linux-amd64-v0.1.0.so.sha256
> ```

### Quick Start

Minimal steps (manual static lib directory):

```bash
# 1) Pull the Go module (latest)
go get github.com/isesword/go-polars@latest

# 2) Prepare the static lib (libpolars_go.a or libpolars_go.lib) and place it in your chosen directory
#    You can download from the corresponding Release asset or build via build.sh

# 3) Tell the compiler to search that directory (adjust the path to your own)
export CGO_LDFLAGS="-L/path/to/your/libss"

# 4) Build / test
go run .
```

> Tip: If you use a precompiled asset, just rename `libpolars_go-<OS>-<ARCH>.a` to `libpolars_go.a` (or `.lib` on Windows) .

### Example

```go
package main

import (
    "fmt"
    "github.com/isesword/go-polars/polars"
)

func main() {
    df, err := polars.ReadCSV("data.csv")
    if err != nil {
        panic(err)
    }
    fmt.Println(df.String())
}
```

### Pre-compiled Binaries

✅ **Available for**:
- Linux x86_64
- macOS x86_64 and ARM64
- Windows x86_64

### Alternative: Build from Source

If pre-compiled binaries aren't available for your platform:

**Prerequisites**:
- **Rust**: Install from [rustup.rs](https://rustup.rs/)
- **Build tools**: `build-essential` (Ubuntu) or equivalent

```bash
git clone https://github.com/isesword/go-polars
cd go-polars
./build.sh
```

## ✨ Features (implemented)

- DataFrame creation and I/O: CSV/JSON readers, writers, Excel via excelize.
- Expressions: arithmetic, comparison, logical, when/otherwise, string ops, null handling, casting.
- GroupBy and aggregations: count, sum/mean/min/max/std/median/quantile, custom Agg.
- Joins: join helpers in `examples/join`, tested in `tests/join_test.go`.
- Sorting & ordering: see `examples/sorting` and `tests/sorting_test.go`.
- Lazy API coverage via expressions and examples (`examples/lazy`).
- Go bindings around core Polars DataFrame operations; see `examples/*` and `tests/*` for usage.

## 🚀 Examples & Quick Start

### Basic Example
Get started with simple DataFrame operations:
```bash
make run-basic-example
```

### Expression Example
Run the full-featured example with complex operations:
```bash
make run-expressions-example
```

### GroupBy Example
Run the GroupBy and aggregation operations demo:
```bash
make run-groupby-example
```

### Available Make Commands
- `make local-build` - Build the library from source (smart build)
- `make force-build` - Force rebuild even if up to date
- `make quick-build` - Smart build (only rebuilds if needed)
- `make run-basic-example` - Run basic DataFrame demo
- `make run-expressions-example` - Run expression operations demo
- `make run-groupby-example` - Run GroupBy and aggregation demo
- `make run-all-examples` - Run all examples

## 🧪 Testing

```bash
# Run all tests
make test

# Quick test run
make test-short

# Test specific functionality
make test-groupby

# Performance benchmarks
make test-bench

# Generate coverage report
make test-coverage

# View coverage in browser
make view-coverage

# Development cycle (quick build + short tests)
make dev
```
## 📋 To do

- [x] Join operations
- [x] Data type conversions: `Cast()`
- [x] Schema inspection
- [x] Null handling: `IsNull()`, `IsNotNull()`, `FillNull()`
- [x] Advanced Aggregations: `Median()`, `Quantile()`, `Var()`, `NUnique()`, `ApproxNUnique()`, `Product()`, `First()`, `Last()`
- [ ] Window functions
- [ ] Pivot & Reshape options
- [x] Additional I/O Formats: `ReadJSON()`, `WriteJSON()`,...
- [x] When/Otherwise logic
- [x] Data Quality & Validation: `IsEmpty()`,...

## 🤝 Contributing

1. Fork the repository
2. Build locally: `./build.sh`
3. Test your changes: `make test`
4. Submit a pull request

## 📄 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
