# 发行与 version.json

## 标签发行与 CI

推送 **`v*`** 标签时，[`.github/workflows/release-lensdns.yml`](../../.github/workflows/release-lensdns.yml) 会发布构建；其中 **`version.json`** 的下载与公告 URL 默认指向 **本仓库的 Releases**。**Release 二进制**默认从 **`…/releases/latest/download/version.json`** 检查更新。

## Docker 镜像

Docker 镜像 **不会** 烘焙上述更新索引——请通过 **`docker pull`** 升级（镜像名见 workflow 默认值，通常为 `ghcr.io/lensdns/lensdns`）。

## 本地 build-release

在未设置 `VERSION_DOWNLOAD_URL`、`ANNOUNCEMENT_URL`、`UPDATE_INDEX_URL` 的情况下本地执行 `make build-release`，行为与 **上游脚本默认**一致（按渠道使用 `static.adtidy.org` 风格的公告/下载基址）。

- **`version.txt`** 仅包含 `version=<tag>`；URL 在 **`version.json`** 中。

## CI 环境变量

CI 默认为 fork 发行设置 `VERSION_DOWNLOAD_URL`、`ANNOUNCEMENT_URL`、`UPDATE_INDEX_URL`。自建流水线请参考 [`scripts/make/build-release.sh`](../../scripts/make/build-release.sh) 与 [`scripts/make/go-build.sh`](../../scripts/make/go-build.sh)。

若将 workflow 复制到其他仓库，请更新其中的 `github.repository` 约束。

GitHub 上至少需要有一次 **Release**，`latest` 相关 URL 才能稳定解析。

打包 `version.json` 中的典型形态：

- `announcement_url` → `https://github.com/LensDNS/LensDNS/releases/tag/<tag>`
- `download_*` → `https://github.com/LensDNS/LensDNS/releases/download/<tag>/<filename>`

**English:** [Releases and version.json](../releases.md)
