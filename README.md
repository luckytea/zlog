# zlog

Simple builder for using the zerologger.

## Использование

```golang
var version = "dev"

log := zlog.New(version)
```

Функиция принимает в себя один аргумент string, в зависимости от него меняется формат вывода данных.

| arg          | description     |
|--------------|-----------------|
| dev          | pretty logging  |
| test         | without logging |
| all the rest | json logging    |
