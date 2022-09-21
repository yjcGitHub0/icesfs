### 存储引擎接口

对于任何其他的存储引擎，其接口通过组合逻辑需要满足 `interface StorageEngine (storage_engine/storage_engine.go)`

对于存储引擎和文件系统，可以通过多种协议进行交互，如 `http`, `grpc`等等

对于新增加的存储引擎，存储引擎出现了错误需要在 `errors/errors.go` 下面新增对应的错误类型进行返回

对于存储引擎接口，需要维护每个文件的链接数，具体维护方式可以是存储引擎自己维护，也可以在文件系统当中利用数据库进行维护

对于任何的异常状态，应该在日志当中进行记录，记录方式可以参考基于 `seaweedfs`实现的存储引擎接口以及 `log/log.go`

总共有如下的接口：

#### 上传文件

```
    PutObject(ctx context.Context, size uint64, file io.Reader, fileName string, compress bool, hosts ...string) (string, error)
```

其中参数为

|    参数    |    类型     |                                    说明                                     |
|:--------:|:---------:|:-------------------------------------------------------------------------:|
|   size   |  uint64   |                传入文件的大小，单位为byte，用于提前告知存储引擎传输文件的大小，从而更好的进行分配                |
|   file   | io.Reader |                                 实际文件的io读流                                 |
| fileName |  string   |                      文件的名称，用于从存储引擎下载时可以直接附带对应的文件名称信息                      |
| compress |   bool    | 表示是否从文件系统端到存储引擎端需要通过压缩的方式传输，其保证了从file是一个完整的未经过压缩的文件，需要在存储引擎接口进行压缩处理后再进行上传 |
|  hosts   | []string  |        表示上传到哪几台主机当中，如果为空则表示上传到任意一台主机均可，如果长度不为1则表示上传到hosts中任意一台主机均可        |

其返回一个 `fid` (类型:  `string`)， `fid`的格式任意，在存储引擎端自己进行解析，但需要保证任何文件拥有全局唯一的 `fid`

#### 删除文件

```
    DeleteObject(ctx context.Context, fid string) error
```

其中参数为

| 参数  |    类型     |                       说明                       |
|:---:|:---------:|:----------------------------------------------:|
| fid |  string   | 对应fid的对象删除，实际上应该为删除一次这个文件链接数，当链接数完全为0时才删除实际的数据 |

对于删除文件，实际上是会出现删除到链接数小于0的情况，此时应该忽略，为正常情况

#### 获得文件的URL

```
    GetFidUrl(ctx context.Context, fid string) (string, error)
```

其中参数为

| 参数  |    类型    |                          说明                          |
|:---:|:--------:|:----------------------------------------------------:|
| fid |  string  | 请求获取文件URL的fid，正常情况下该fid下的文件还并未被删除，如果已经被删除则表示文件系统出现问题 |

其返回一个 `URL` (类型:  `string`)，用于给用户直接从存储引擎当中获取数据

#### 获得存储引擎下的所有主机

```
    GetHosts(ctx context.Context) ([]string, error)
```

获得所有存储引擎下控制的主机的`host`

#### 增加文件的链接数

```
    AddLink(ctx context.Context, fid string) error
```

其中参数为

| 参数  |    类型    |                        说明                        |
|:---:|:--------:|:------------------------------------------------:|
| fid |  string  | 增加链接数的fid，正常情况下该fid下的文件还并未被删除，如果已经被删除则表示文件系统出现问题 |