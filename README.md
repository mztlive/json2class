# classgenerator
### 用途
根据传入的json生成php对象
### 用法
```bash
Usage of ./classgenerator:
  -f string
        Save Folder (default "./") #类文件要存放的地方
  -json string
        json string #json内容（不要格式化）
  -namespace string
        PHP Class Namespace (default "Example") #类的命名空间


调用例子：
    ./classgenerator --namespace="RequestExample" \
    --f="../src/RequestExample" \
    --json='{"hello":"world", "number":[1,2,3], "objectList":[{"id": 1}], "obj": {"name": "test", "age": 10}}'
