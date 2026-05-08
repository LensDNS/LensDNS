# DNS 引擎（LensDNS dnsproxy）

LensDNS 将上游 [`github.com/AdguardTeam/dnsproxy`](https://github.com/AdguardTeam/dnsproxy) 替换为 [`github.com/LensDNS/dnsproxy`](https://github.com/LensDNS/dnsproxy)。因此 **DNS 代理核心** 可能与原版 AdGuard Home 存在差异，包括但不限于：

- 面向 DoT 的 TLS 超时行为
- TCP Keep-Alive
- DoT 监听端的 TFO（取决于平台）
- TLS 会话恢复
- RST/关闭路径加固
- 对 **DNS-over-TCP 两字节长度前缀** 的更稳健读取（减轻短读导致的错误组帧）

**开关、默认值与 CLI** 以 **dnsproxy** 仓库文档为准。

AdGuard Home 通过 **`dns.*` YAML / Web UI** 暴露与集成相关的设置（例如 PPv2 与 `trusted_proxies`）。

详见：[github.com/LensDNS/dnsproxy](https://github.com/LensDNS/dnsproxy)。

**English:** [DNS engine](../engine.md)
