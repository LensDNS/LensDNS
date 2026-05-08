# LensDNS 文档（中文）

本目录收录面向 [LensDNS/LensDNS](https://github.com/LensDNS/LensDNS)（基于 [AdguardTeam/AdGuardHome](https://github.com/AdguardTeam/AdGuardHome) 的发行版）的 **运维与架构说明**。

上游通用的故事、构建说明与社区资源见仓库根目录 [README](../../README_zh.md)（内含嵌入的上游 README）。

**English:** [docs index in English](../README.md)

## 品牌资源（源 SVG，默认不打进前端产物）

LensDNS 长方形 Logo 与正方形 favicon 与上游素材并列放在 [`doc/`](../../doc/)：

- [doc/lensdns-logo.svg](../../doc/lensdns-logo.svg) — 横向 Logo（README 用）
- [doc/favicon.svg](../../doc/favicon.svg) — 方形图标，适合低分辨率 / favicon

若要在 Web UI 中一并分发，请在变更中显式复制或引用到 `client/public/assets/`（并按需添加 `<link rel="icon" …>`）。

## 指南

- [Proxy Protocol v2（PPv2）](ppv2.md) — 负载均衡拓扑、严格语义、可信代理、YAML 示例。
- [发行与 version.json](releases.md) — GitHub Actions 发行、`version.json`、Docker 与二进制更新差异。
- [DNS 引擎](engine.md) — [`LensDNS/dnsproxy`](https://github.com/LensDNS/dnsproxy) 集成与行为差异。

## 相关仓库

- DNS 代理核心：[github.com/LensDNS/dnsproxy](https://github.com/LensDNS/dnsproxy)
