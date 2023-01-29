```shell
# 生成model
gentool -db=mysql -dsn="root:123456@tcp(127.0.0.1:3307)/yangtze-admin?charset=utf8mb4&parseTime=True&loc=Local" -modelPkgName=entity -outPath="./internal/model/sysUser" -tables=sys_user
```