# media vault

离线动画资源检索工具。

## 预览



## 开发环境

- go1.23.4
- nodejs v22.12.0
- cnpm v9.4.0

前端框架 vite + vue + element-plus
后端框架 gin + sqlite

封面抽帧需要安装 ffmpeg。

## 目录结构

前端项目保存在 media-vault 下。
后端项目大体参考 [project layout](https://github.com/golang-standards/project-layout)。

API 简便起见大量使用 POST，没有遵循 RESTful 设计。

## 参与开发

还没有做好协作开发的准备，代码规范就是先仿照已有代码风格，其它随便，我看不懂就不合并。

commit 信息参考 [conventional commit message](https://www.conventionalcommits.org/en/v1.0.0/)

提交 PR 尽可能小且只解决单个明确的问题/需求，以减轻 review 压力。

开始写代码前先提 issue 讨论是否属于问题/需求，以及是否正在开发中，以免重复开发浪费精力。

## LICENSE

GPL v2
