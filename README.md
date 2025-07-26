# SCMZU-Party-Building
中南民族大学计算机学院（人工智能学院）2023级本科生党支部智慧党建项目

本次项目目录定义如下

```
├── api/app/application // App层，处理Adapter层适配过后与框架、协议等无关的业务逻辑
├   ├── api 处理OpenAPI 接口请求
├   ├── web  请求Web页面请求
├   ├── scheduler/task  //处理定时任务，比如Cron格式的定时Job
├── domain // Domain层，最纯粹的业务实体及其规则的抽象定义
│   ├── interface/gateway // 领域网关，model的核心逻辑以Interface形式在此定义，交由biz层去实现
│   └── model // 领域模型实体
├── biz/module 业务/业务模块层
├   └──module1
├        ├──dao 数据库层
├        ├──model 业务模型
├        ├──entity 数据库模型
├        ├──service 业务逻辑
├── configs 配置
├── init    //系统初始化
├── cmd    // 存放可执行程序，main 包放这个目录中
├── public  // 静态公共资源，实际项目会将其存入 CDN
├── test // 单元测试之外的测试程序、测试数据
├── util/tools 工具包
├── main.go 项目运行入口
└── pkg // 各层可共享的公共组件代码
```



