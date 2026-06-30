# template.go.cobra

Go + `Cobra` + `Viper` 用の CLI テンプレート

## 初回

- `__REPLACE_MODULE_NAME__` を置換
- 必要なら `__ORGANIZE_NAME__` も置換
- `go.mod`, `main.go`, `cmd/root.go`, `Taskfile.yml` を確認
- `go mod tidy`
- `cmd/sample.go` は消すか作り変える

## メモ

- `cmd/root.go` で設定ファイルを読む
- `cmd/version.go` で build info を出す
- `Taskfile.yml` で build / test / lint / run をまとめる
- GitHub Actions は `test`, `build`, `golangci-lint` 用

## よく使うコマンド

```bash
go mod tidy
go run . version
go run . sample
task build
task test
task lint
task run CLI_ARGS="version"
```

## Task

| コマンド | 用途 |
| --- | --- |
| `task` | タスク一覧 |
| `task init MODULE_NAME=...` | `go.mod` の module 更新 |
| `task mod-deps` | 依存整理 |
| `task run CLI_ARGS="..."` | `go run .` |
| `task build` | ビルド |
| `task build:prod` | 最適化ビルド |
| `task exe CLI_ARGS="..."` | ビルドして実行 |
| `task test` | テスト |
| `task test-coverage` | カバレッジ付きテスト |
| `task lint` | lint |
| `task all` | lint + test + build |
| `task clean` | 成果物削除 |

`task init` は `go.mod` しか触らない。
プレースホルダーの一括置換は別でやる。

## Cobra

```bash
go install github.com/spf13/cobra-cli@latest
cobra-cli add your-command
```

## Config

- `--config <path>` で明示指定可能
- 未指定時は XDG 系の設定ディレクトリを探索
- Windows では `%USERPROFILE%\.config\<app-name>` も探索

## 構成

```text
.
├── cmd/
│   ├── root.go
│   ├── sample.go
│   └── version.go
├── .github/workflows/
├── .dev/task/
├── main.go
├── go.mod
├── Taskfile.yml
└── README.md
```
