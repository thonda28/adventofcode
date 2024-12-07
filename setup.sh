#!/bin/bash

WORKING_DIR=$(dirname "$(realpath "$0")")

# 引数のチェック
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 yyyy dd"
    exit 1
fi

# 入力引数
year=$1
day=$2

# yyyy が4桁の数字であることを確認
if ! [[ $year =~ ^[0-9]{4}$ ]]; then
    echo "Error: yyyy must be a 4-digit number."
    exit 1
fi

# dd が1～2桁の数字であることを確認
if ! [[ $day =~ ^[0-9]{1,2}$ ]]; then
    echo "Error: dd must be a 1- or 2-digit number."
    exit 1
fi

# dd を0埋め（1桁の場合は先頭に0を追加）
day=$(printf "%02d" $day)

# ディレクトリパス
target_dir="$WORKING_DIR/$year/$day"

# ディレクトリの作成（既存の場合はスキップ）
if [ ! -d "$target_dir" ]; then
    mkdir -p "$target_dir"
    echo "Created directory: $target_dir"
else
    echo "Directory already exists: $target_dir"
fi

# go mod init を実行（ディレクトリ内で実行）
(
    cd "$target_dir" || exit
    if [ ! -f "go.mod" ]; then
        go mod init "github.com/thonda28/adventofcode/$year/$day"
        echo "Initialized Go module: github.com/thonda28/adventofcode/$year/$day"
    else
        echo "Go module already initialized in: $target_dir"
    fi
)

# テンプレートファイルのコピー
template_src="$WORKING_DIR/template/dayxx.go"
template_dest="$target_dir/day$day.go"
if [ ! -f "$template_dest" ]; then
    if [ -f "$template_src" ]; then
        cp "$template_src" "$template_dest"
        echo "Copied template to: $template_dest"
    else
        echo "Error: Template file not found at $template_src"
        exit 1
    fi
else
    echo "Template file already exists: $template_dest"
fi

# 入力ファイルの作成
for file in "day$day.example" "day$day.input"; do
    target_file="$target_dir/$file"
    if [ ! -f "$target_file" ]; then
        touch "$target_file"
        echo "Created file: $target_file"
    else
        echo "File already exists: $target_file"
    fi
done

echo "Setup complete!"
