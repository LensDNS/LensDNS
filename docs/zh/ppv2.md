# Proxy Protocol v2（PPv2）

DNS 常部署在 L4/L7 负载均衡之后。若在传输层拿不到真实客户端身份，AdGuard Home 只能看到负载均衡地址，导致按客户端限速、滥用归因与可观测性变差。LensDNS 在 **DNS-over-TCP** 与 **DoT** 监听上实现 **Proxy Protocol v2**，在 **LB → AdGuard Home** 这一段恢复真实的客户端语义。

YAML 键：

- `dns.tcp-proxy-protocol-v2`
- `dns.tls-proxy-protocol-v2`
- `dns.trusted_proxies`

严格开/关语义；DoT 顺序：**PPv2 → TLS → DNS 组帧**。

## 严格语义

- 在 **启用** PPv2 的监听上，连接 **必须** 携带合法 PPv2 头，否则被拒绝。
- 在 **关闭** PPv2 的监听上，客户端 **不得** 发送 PPv2（避免直连与代理混用）。
- **DoT**：在 TLS 握手 **之前** 解析 PPv2。顺序：PPv2 → TLS → DNS-over-TCP 组帧。

## 信任边界

- **仅**将可信的 LB/代理 CIDR 写入 `dns.trusted_proxies`。
- 通过网络策略（ACL/SG）限制 **只有** 这些代理能访问启用 PPv2 的监听。
- **不要**在公网边缘轻信客户端发来的 PPv2（例如粗心地 `accept-proxy`），否则可能出现源地址伪造。

## 拓扑（示意）

关闭 PPv2：

```text
+--------+      +-------------------+      +------------------+
| Client | ---> | HAProxy/Nginx/LB  | ---> | AdGuard Home     |
+--------+      +-------------------+      +------------------+
```

说明：AdGuard Home 看到的 `remote addr` 为 LB 地址。

启用 PPv2：

```text
+--------+      +-------------------+   send-proxy-v2 / PPv2   +------------------+      +--------------+
| Client | ---> | HAProxy/Nginx/LB  | -----------------------> | AdGuard Home     | ---> | Upstream DNS |
+--------+      +-------------------+     (require PPv2)       +------------------+      +--------------+
```

说明：Client → LB **不**携带 PPv2；PPv2 由 LB 在 **LB → AdGuard Home** 的后端连接上注入。

## YAML（AdGuard Home）

```yaml
dns:
  tcp-proxy-protocol-v2: true
  tls-proxy-protocol-v2: true
  trusted_proxies:
    - 10.0.0.0/8
```

## 反向代理

Nginx stream 端到端使用 `proxy_protocol`；HAProxy 对 AdGuard Home 使用 `send-proxy-v2`（DoT 常为 TLS 透传）；Envoy 在适用场景下于 TLS 之前转发 PPv2。

## HAProxy 最小片段（DoT → 本机 AdGuard Home）

```text
backend bk_adguard_dot
  server adguardhome 127.0.0.1:853 send-proxy-v2
```

然后启用 `dns.tls-proxy-protocol-v2`（按需启用 `dns.tcp-proxy-protocol-v2`）。完整语法请以 HAProxy 官方文档为准。

## DoQ / UDP

QUIC/UDP 语义与 TCP/DoT 不同。**本文 PPv2 仅适用于 TCP/DoT。** 更广的讨论见 [LensDNS/dnsproxy](https://github.com/LensDNS/dnsproxy) Discussions。

**English:** [Proxy Protocol v2](../ppv2.md)
