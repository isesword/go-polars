# go-polars

<p align="center">
    <img src="docs/assets/images/go-rust.png" width="300"/>
</p>

本项目为 Polars 数据处理库提供 Go 绑定。

## 🐻‍❄️ 什么是 Polars？
Polars 是一个开源的数据处理库，在单机场景下以高性能著称，提供结构化且类型安全的 API，表达力强且易于使用。

https://github.com/pola-rs/polars

## 📦 安装

> [!NOTE]
> **构建与安全注意事项**
>
> 由于 GitHub Actions 资源受限，Polars 的 Rust 绑定目前在本地机器编译。为提升透明度和可验证性，我们采取了以下措施：
>
> - **🔍 可复现构建**：所有构建脚本可在 [`./scripts`](./scripts) 查阅
> - **🔐 校验和验证**：每个二进制发布都包含 SHA256 和 MD5 校验
> - **📋 构建透明度**：发布说明包含构建环境与依赖版本
> - **🏗️ 自行编译**：随时可运行 `./build.sh` 从源码构建
>
> **校验下载的二进制示例：**
> ```bash
> sha256sum -c libpolars_go-linux-amd64-v0.1.0.so.sha256
> ```

### 快速开始

最小步骤（手动指定静态库目录）：

```bash
# 1) 拉取 Go 模块
go get github.com/isesword/go-polars@latest

# 2) 准备静态库（libpolars_go.a 或 libpolars_go.lib）并放到自选目录
#    可从对应 Release 资产下载，或用 build.sh 自行构建

# 3) 告诉编译器去该目录找库（将路径替换为你的实际路径）
 export CGO_LDFLAGS="-L/Users/esword/GolandProjects/src/Test/libs"

# 4) 构建 / 测试
 go run .
```

> 提示：如果使用预编译资产，只需将 `libpolars_go-<OS>-<ARCH>.a` 重命名为 `libpolars_go.a`（Windows 为 `.lib`），放入 `POLARS_BIN_DIR` 后即可编译。

### 示例

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

### 预编译二进制

✅ **已提供**：
- Linux x86_64
- macOS x86_64 和 ARM64
- Windows x86_64

### 备选：从源码构建

若你的平台没有预编译包：

**前置要求**：
- **Rust**：安装自 [rustup.rs](https://rustup.rs/)
- **构建工具**：`build-essential`（Ubuntu）或等价工具

```bash
git clone https://github.com/isesword/go-polars
cd go-polars
./build.sh
```

## ✨ 已实现功能

- DataFrame 创建与 I/O：CSV/JSON 读写，Excel（excelize）。
- 表达式：算术、比较、逻辑、when/otherwise、字符串操作、空值处理、类型转换。
- GroupBy 与聚合：count，sum/mean/min/max/std/median/quantile，自定义 Agg。
- Join：`examples/join` 中的示例，`tests/join_test.go` 覆盖。
- 排序与排序规则：见 `examples/sorting` 与 `tests/sorting_test.go`。
- Lazy API：通过 expressions 与 `examples/lazy` 覆盖核心用法。
- Go 绑定封装核心 Polars DataFrame 操作，可参考 `examples/*` 与 `tests/*`。

## 🚀 示例与快捷入口

### Basic Example
```bash
make run-basic-example
```

### Expression Example
```bash
make run-expressions-example
```

### GroupBy Example
```bash
make run-groupby-example
```

### 可用的 Make 命令
- `make local-build` - 从源码智能构建
- `make force-build` - 即使最新也强制重建
- `make quick-build` - 智能构建（仅需时重建）
- `make run-basic-example` - 运行基础 DataFrame 示例
- `make run-expressions-example` - 运行表达式示例
- `make run-groupby-example` - 运行分组聚合示例
- `make run-all-examples` - 运行全部示例

## 🧪 测试

```bash
# 运行全部测试
make test

# 快速测试
make test-short

# 指定模块
make test-groupby

# 基准
make test-bench

# 覆盖率
make test-coverage

# 浏览器查看覆盖率
make view-coverage

# 开发循环（快构建 + 短测）
make dev
```

## 📋 待办

- [x] Join
- [x] 类型转换：`Cast()`
- [x] Schema 检查
- [x] 空值处理：`IsNull()`、`IsNotNull()`、`FillNull()`
- [x] 高级聚合：`Median()`、`Quantile()`、`Var()`、`NUnique()`、`ApproxNUnique()`、`Product()`、`First()`、`Last()`
- [ ] 窗口函数
- [ ] Pivot & Reshape
- [x] 更多 I/O：`ReadJSON()`、`WriteJSON()` 等
- [x] When/Otherwise 逻辑
- [x] 数据质量校验：`IsEmpty()` 等

## 🤝 参与贡献

1. Fork 本仓库
2. 本地构建：`./build.sh`
3. 运行测试：`make test`
4. 提交 PR

## 📄 许可证

基于 MIT 许可证，详见 [LICENSE](LICENSE)。
