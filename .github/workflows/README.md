# Workflows ガイド

このディレクトリには、CI とリリース用ワークフローが含まれます。

## CI ワークフロー

- 対象ファイル: `ci.yml`、`golangci-lint.yml`
- 主な実行内容:
  1. `go fmt`
  2. `go vet`
  3. `go test`
  4. `go build`

### ローカル確認

CI 実行前は以下のタスクで検証できます。

```bash
task lint
task test
task build
```

または一括実行:

```bash
task all
```

## release ワークフロー

release は sample を用途に応じてリネームして使います。

### ファイル構成

- `gh-release.yml.sample`: GitHub Releases 用 sample
- `fj-release.yml.sample`: Forgejo Releases 用 sample

### 使い方

1. GitHub で使う場合
   - `gh-release.yml.sample` を `gh-release.yml` にリネームして有効化
2. Forgejo で使う場合
   - `fj-release.yml.sample` を `fj-release.yml` にリネームして有効化
3. GitHub と Forgejo の両方で使う場合
   - GitHub 側: `.github/workflows/*.yml`
   - Forgejo 側: `.forgejp/workflows/*.yml` にコピーして分離

## 実行状況の確認

- GitHub: リポジトリの **Actions** タブ
- Forgejo: 対象リポジトリの Actions / CI 画面
