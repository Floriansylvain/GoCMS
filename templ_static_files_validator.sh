#!/bin/bash

find_html_files() {
    local directory="$1"
    find "$directory" -type f -name "*.html"
}

extract_referenced_files() {
    local html_file="$1"
    grep -oP '(?<=src="/static/|href="/static/)[^"]+' "$html_file"
}

is_link() {
    local path="$1"
    [[ "$path" =~ ^https?://.* ]]
}

validate_referenced_files() {
    local source_dir="$1"
    local target_dir="$2"

    local html_files
    html_files=$(find_html_files "$source_dir")

    for html_file in $html_files; do
        local referenced_files
        referenced_files=$(extract_referenced_files "$html_file")

        for referenced_file in $referenced_files; do
            referenced_file="${referenced_file#/}"

            if is_link "$referenced_file"; then
                continue
            fi

            local full_referenced_path="$target_dir$referenced_file"

            if [ ! -e "$full_referenced_path" ]; then
                echo "Error: $full_referenced_path not found (referenced in $html_file)"
                exit 1
            fi
        done
    done
}

source_dir="./adapters/secondary/gateways/web/templates/"
target_dir="./api/static/"

validate_referenced_files "$source_dir" "$target_dir"

echo "All files referenced in HTML files are present in $target_dir"
