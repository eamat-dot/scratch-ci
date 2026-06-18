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

1. GitHub のみで使う場合
   - `gh-release.yml.sample` を `gh-release.yml` にリネームして有効化

2. Forgejo のみで使う場合
   - `fj-release.yml.sample` を `fj-release.yml` にリネームして有効化
   - フォルダは `.github/workflows/` のままでよい

3. GitHub と Forgejo の両方で使う場合
   - `.github/workflows/` と `.forgejo/workflows/` に分離する
   - GitHub 側: `.github/workflows/*.yml` (`gh-release.yml` を有効化)
   - Forgejo 側: `.forgejp/workflows/*.yml` (`.github/workflows/`から全てコピーし `fj-release.yml` を有効化)


#### Forgejo ワークフローの優先順位
> **NOTE:**
> - `.forgejo/workflows/` → `.gitea/workflows/` → `.github/workflows/` の順
> - フォールバックは「全か無か」
> - `.forgejo/workflows/` にファイルが1つでもあれば `.github/workflows/` は**完全に無視**され、両方を同時に動かすことはできない  \
> - `.forgejo/workflows/` ディレクトリが存在しない場合は、`.github/workflows/` が使われる

## 実行状況の確認

- GitHub: リポジトリの **Actions** タブ
- Forgejo: 対象リポジトリの Actions / CI 画面
