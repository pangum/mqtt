# 数据库

数据库在大家的开发经常都会使用到，所以`盘古`也做了相应的适配，包括

- 关系型数据库`xorm`
- Elasticsearch
- Redis

之所以把这三者混在一起，是因为现在`盘古`势力还很弱小，等将来功能强大后，可以考虑分开，参考`spring-boot-starter`的方法来编写插件

## XORM数据库操作

关系型数据库用得最多，在这方面，`盘古`支持以下几种功能

- 快速集成数据库操作
- 更友好的事务支持

### 快速集成

集成只需要三个步骤

- 引入数据库包（因为`盘古`支持数据迁移，内置了`xorm`，所以引入不引入都不重要）
- 配置数据库连接配置
- 依赖数据库操作对象

## Elasticsearch

TODO

## Redis

TODO
