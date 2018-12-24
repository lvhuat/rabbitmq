rabbit
------------
这是对[amqp](github.com/streadway/amqp)库的常见的使用方法的封装。

注意
-----
1. 在使用本工具之前，需要充分的amqp协议知识
2. 包内不处理重连逻辑(也不建议)，请在业务层面做相应处理

关于setting
-----
使用AMQP时设置是很多的，为了灵活并且简洁，使用了map[string]bool的方式来传递设置,以下是默认配置：
+ exchange
```
// durable=true,autodelete=false 始终保持
// durable=true,autodelete=true 服务启动时无绑定关系即删除
// durable=false,autodelete=true 服务启动时无绑定关系即删除
// durable=false,autodelete=false 服务启动时删除
"exchange/durable":    true
"exchange/autodelete": false
// 被设置后，不能用于接受消息，但是你可以用它来建立你的内部消息拓扑结构（二级exchange）
"exchange/internal":  false
```
+ queue 
```
// durable=true,autodelete=false 始终保持，仅能绑定durable的exchange
// durable=true,autodelete=true 服务启动时无绑定关系即删除，仅能绑定durable的exchange
// durable=false,autodelete=true 无绑定关系后一会后即删除，仅能绑定非durable的exchange
// durable=false,autodelete=false 服务启动时删除,仅能绑定非durable的exchange
"queue/durabale":   true
"queue/autodelete": false
// 排他队列，同连接可访问，其他连接不可见，名字唯一，连接释放删除(durabale被忽略)
"queue/exclusive":  false 
``
```
+ consume
```
// 在接受到消息以后理解反馈给服务器，注意模块并不理会业务是否处理得过来
"consume/autoack":   true
// 排他消费，绑定时会检查队列是否存在其他消费者，
"consume/exclusive": false
// RabbitMQ不支持此设置
"consume/nolocal":   false
```

使用方法
-------
参考example文件夹
