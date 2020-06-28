# 极客大学「算法训练营-第 9 期」作业提交仓库

## 讲师课件下载地址

请大家通过该链接查看讲师课件并进行下载，链接:https://pan.baidu.com/s/1VQEJb6BE1YL4AbEZT0icYg 密码:un6x

## 仓库目录结构说明

1. `week01/` 代表第一周作业提交目录，以此类推。
2. 请在对应周的目录下新建或修改自己的代码作业。
3. 每周均有一个 `REDAME.md` 文档，你可以将自己当周的学习心得以及做题过程中的思考记录在该文档中。

## 作业提交规则

1. 先将本仓库 Fork 到自己 GitHub 账号下。
2. 将 Fork 后的仓库 Clone 到本地，然后在本地仓库中对应周的目录下新建或修改自己的代码作业，当周的学习总结写在对应周的 README.md 文件里。
3. 在本地仓库完成作业后，push 到自己的 GitHub 远程仓库。
4. 最后将远程仓库中当周的作业链接，按格式贴到班级仓库对应学习周的 issue 下面。
5. 提交 issue 请务必按照规定格式进行提交，否则作业统计工具将抓取不到你的作业提交记录。

详细的作业提交流程可以查阅：https://shimo.im/docs/m5rtM8K8rNsjw5jk/

## 注意事项

如果对 Git 和 GitHub 不太了解，请参考 [Git 官方文档](https://git-scm.com/book/zh/v2) 或者极客时间的[《玩转 Git 三剑客》](https://time.geekbang.org/course/intro/145)视频课程。

######################

1. harbor 报 500，一般是内部配置错误
2. 重启 adminserver 模块重启后恢复
3. 分析发现 adminserver 重启后，数据库的 properties 表被更新，这些都是 harbor 的各个组件的连接配置
4. 最终导致了新的 harbor 的 adminserver 启动后覆盖了老的 harbor 的配置，最终导致 core 模块链接不到 Registry 模块，最后获取不到 docker 镜像的 tag。

Q: 为什么在高峰期做部署变更
A: Harbor 没有变更 SOP 可寻，部署 HA 在测试（没有覆盖到数据库）对线上数据无感。增加 harbor 变更的 SOP 和对数据变更的测试。

Q: 变更后为什么没有及时测试现有的 hub 问题
A: 测试计划是先测试新的 harbor 实例，再测试线上的 harbor。测试顺序有问题，这个问题在增加到 SOP 里。

Q: 为什么没有报警出来？
A: harbor 的域名报警云平台基础设施没有订阅，sre 同学是有收到告警的

Q: 为什么第一时间怀疑数据库问题？
A: 登录到 harbor 上， 发现所有个 repo 的 tag 全部是 0， 怀疑是数据库数据丢失问题。通过查看数据库的数据，发现 tag 信息不再数据库里，是 Registry 服务自己维护的（冷静下来其实之前测试的过程中  知道 tag 信息是 Registry 服务维护的，不在数据库里维护）

Q: 为什么第一时间没有看日志，而是怀疑数据库问题

A: Harbor Web 界面访问正常，只有 repo 的 tag 有异常，而且老的 harbor 实例没有做任何变更，所以一开始没有把问题放到 harbor 的日志上。

Q: 为什么重启老 harbor 的 adminServer 服务就好了
A: adminserver 会把本地的配置重新覆盖到数据库中，恢复到了之前的状态（各个组件连接正常）

Q: 是否有回退方案
A: 直接停止新的 harbor 实例，回退方案比较糙，后续完善回退方案。

kubectl drain --grace-period=300 --ignore-daemonsets=true --delete-local-data=true g1-med-online1-84 --dry-run
kubectl drain --grace-period=300 --ignore-daemonsets=true --delete-local-data=true g1-med-online1-84
