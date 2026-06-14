# Forgejo Actions CI

Forgejoは `.forgejo/workflows/` → `.gitea/workflows/` → `.github/workflows/` の順に確認します

フォールバックは「全か無か」です。`.forgejo/workflows/` にファイルが1つでもあれば `.github/workflows/` は**完全に無視**されます。両方を同時に動かすことはできません。

逆に `.forgejo/workflows/` ディレクトリが存在しない場合は、`.github/workflows/` が使われます。
