## split Tools
    
```bash
➜  go-split-file git:(master) ./split -h
NAME:
   split - file split

USAGE:
   split [global options] command [command options] [arguments...]

VERSION:
   0.0.0

AUTHOR:
   Zhang Jian Xin

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value  filename
   --line value, -l value  split lines (default: 0)
   --buff value, -b value  split buffer (default: 0)
   --help, -h              show help
   --version, -v           print the version


```


* 使用

```bash
# 编译
go build -ldflags '-w -s' -o split ./
# 运行
./split -f data.csv -b 10 #将文件以10mb 一个part 分割
./split -f data.csv -l 10000 #将文件以10000行一个part 分割（默认会加上Head）

```

#### 超大文件采用 buffer 分割，文本文件请采用Line分割


