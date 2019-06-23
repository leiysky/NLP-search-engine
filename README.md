# Simple Search Engine

使用 golang 编写的简易搜索引擎。

使用 Jieba 进行分词，实现了倒排索引和基于 tf-idf 的文档排序。

# Usage

## Build 

```shell
$ go get -u && go build
```

## Run

```shell
$ ./search-engine
```

## API

HTTP API:
```
GET localhost/search?q=${your_query}
```