# Proxy Protocol v2 (PPv2)

DNS is often placed behind L4/L7 load balancers. Without transport-level client identity, AdGuard Home only sees the LB address, so per-client rate limiting, abuse attribution, and observability degrade. LensDNS implements **Proxy Protocol v2** on **DNS-over-TCP** and **DoT** listeners to restore real client semantics on the **LB → AdGuard Home** hop.

YAML keys:

- `dns.tcp-proxy-protocol-v2`
- `dns.tls-proxy-protocol-v2`
- `dns.trusted_proxies`

Strict on/off semantics; DoT order: **PPv2 → TLS → DNS framing**.

## Strict semantics

- On listeners where PPv2 is **enabled**, connections **must** carry a valid PPv2 header or they are rejected.
- On listeners where PPv2 is **disabled**, clients **must not** send PPv2 (avoids mixed direct/proxy behaviour).
- **DoT**: PPv2 is parsed **before** the TLS handshake. Order: PPv2 → TLS → DNS-over-TCP framing.

## Trust boundary

- Put **only** trusted LB/proxy CIDRs in `dns.trusted_proxies`.
- Restrict the network (ACL/SG) so **only** those proxies can reach PPv2-enabled listeners.
- **Do not** accept client-supplied PPv2 on a public edge (e.g. careless `accept-proxy`), or source-address spoofing becomes possible.

## Topology (sketch)

PPv2 disabled:

```text
+--------+      +-------------------+      +------------------+
| Client | ---> | HAProxy/Nginx/LB  | ---> | AdGuard Home     |
+--------+      +-------------------+      +------------------+
```

Note: AdGuard Home sees `remote addr` = LB address.

PPv2 enabled:

```text
+--------+      +-------------------+   send-proxy-v2 / PPv2   +------------------+      +--------------+
| Client | ---> | HAProxy/Nginx/LB  | -----------------------> | AdGuard Home     | ---> | Upstream DNS |
+--------+      +-------------------+     (require PPv2)       +------------------+      +--------------+
```

Note: Client → LB does **not** carry PPv2. PPv2 is injected by the LB on the **LB → AdGuard Home** backend connection.

## YAML (AdGuard Home)

```yaml
dns:
  tcp-proxy-protocol-v2: true
  tls-proxy-protocol-v2: true
  trusted_proxies:
    - 10.0.0.0/8
```

## Reverse proxies

Nginx stream uses `proxy_protocol` end-to-end; HAProxy uses `send-proxy-v2` toward AdGuard Home (DoT is often TLS passthrough); Envoy forwards PPv2 before TLS where applicable.

## Minimal HAProxy fragment (DoT → local AdGuard Home)

```text
backend bk_adguard_dot
  server adguardhome 127.0.0.1:853 send-proxy-v2
```

Then enable `dns.tls-proxy-protocol-v2` (and `dns.tcp-proxy-protocol-v2` if needed). Follow HAProxy’s own docs for full syntax.

## DoQ / UDP

QUIC/UDP semantics differ from TCP/DoT. **PPv2 here is TCP/DoT only.** Broader discussion: [LensDNS/dnsproxy](https://github.com/LensDNS/dnsproxy) Discussions.
