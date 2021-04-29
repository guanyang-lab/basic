# basic
##项目需求
自动生成控制层、实现层以及models层

因为我们表生成是采用结构体生成的，所以只提供生成代码，需传入结构体名称指定结构体生成，或者传入数据库生成所有的表的crud。
如果建表是采用的数据建表，可以参考github上的go-mygen，采用从数据库读取字段、备注等生成结构体以及crud代码

go build main.go

windows:  main.exe -s User,Role   或者 main.exe -a 127.0.0.1:3306 -u root -p 123456 -P t_

linux: ./main -s User,Role 或者 ./main -a 127.0.0.1:3306 -u root -p 123456 -P t_
    
