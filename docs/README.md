# LensDNS documentation

This directory holds **LensDNS-specific** operational and architecture notes for [LensDNS/LensDNS](https://github.com/LensDNS/LensDNS) (a distribution based on [AdguardTeam/AdGuardHome](https://github.com/AdguardTeam/AdGuardHome)).

For the generic upstream story, build instructions, and community resources, see the [main README](../README.md) (including the embedded upstream README).

**简体中文：** [docs/zh/README.md](zh/README.md)

## Branding (source SVGs, not bundled by default)

LensDNS logo and square favicon live next to upstream artwork under [`doc/`](../doc/):

- [doc/lensdns-logo.svg](../doc/lensdns-logo.svg) — wide logo (README)
- [doc/favicon.svg](../doc/favicon.svg) — square icon for low-resolution / favicon use

To ship them in the web UI, copy or reference from `client/public/assets/` in a deliberate change (and add `<link rel="icon" …>` if needed).

## Guides

- [Proxy Protocol v2 (PPv2)](ppv2.md) — LB topology, strict semantics, trusted proxies, YAML examples.
- [Releases and version.json](releases.md) — GitHub Actions releases, `version.json`, Docker vs binary updates.
- [DNS engine](engine.md) — [`LensDNS/dnsproxy`](https://github.com/LensDNS/dnsproxy) integration and behaviour differences.

## Related repositories

- DNS proxy core: [github.com/LensDNS/dnsproxy](https://github.com/LensDNS/dnsproxy)
