#!/bin/bash
project_script_dir=$(cd $(dirname "$0");pwd)
echo '脚本所在文件夹：' $project_script_dir
project_root_dir=$(dirname "$project_script_dir")
echo '项目所在文件夹：' $project_root_dir
project_build_file="$project_root_dir/build/litkeep"
project_build_file_normalized=$(echo $project_build_file | sed "s/\/\{2,\}/\//g")
echo '编译文件：' $project_build_file_normalized
cd "$project_root_dir"
go build -o ./build/litkeep ./cmd/main.go
