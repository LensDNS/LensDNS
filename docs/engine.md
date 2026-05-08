# DNS engine (LensDNS dnsproxy)

LensDNS replaces upstream [`github.com/AdguardTeam/dnsproxy`](https://github.com/AdguardTeam/dnsproxy) with [`github.com/LensDNS/dnsproxy`](https://github.com/LensDNS/dnsproxy). The **DNS proxy core** may therefore differ from stock AdGuard Home, including:

- DoT-oriented TLS timeout behaviour
- TCP keep-alive
- TFO on DoT listeners (platform-dependent)
- TLS session resumption
- RST/close-path hardening
- Robust reading of the **two-byte DNS-over-TCP length prefix** (reduces mis-framing on short reads)

**Flags, defaults, and CLI** are documented in the **dnsproxy** repository.

AdGuard Home exposes integration-related settings via **`dns.*` YAML / the web UI** (e.g. PPv2 and `trusted_proxies`).

See: [github.com/LensDNS/dnsproxy](https://github.com/LensDNS/dnsproxy).
