# Releases and version.json

## Tag releases and CI

When you push a **`v*`** tag, [`.github/workflows/release-lensdns.yml`](../.github/workflows/release-lensdns.yml) publishes builds whose **`version.json`** download and announcement URLs default to **this repository’s Releases**. **Release binaries** default to checking **`…/releases/latest/download/version.json`**.

## Docker images

Docker images **do not** bake that update index—upgrade with **`docker pull`** (see default image name in the workflow, typically `ghcr.io/lensdns/lensdns`).

## Local build-release

A local `make build-release` **without** `VERSION_DOWNLOAD_URL`, `ANNOUNCEMENT_URL`, and `UPDATE_INDEX_URL` matches **upstream script defaults** (`static.adtidy.org`-style announcement/download bases per channel).

- **`version.txt`** is only `version=<tag>`; URLs live in **`version.json`**.

## CI environment variables

CI sets `VERSION_DOWNLOAD_URL`, `ANNOUNCEMENT_URL`, and `UPDATE_INDEX_URL` by default for fork releases. Custom pipelines: see [`scripts/make/build-release.sh`](../scripts/make/build-release.sh) and [`scripts/make/go-build.sh`](../scripts/make/go-build.sh).

If you copy the workflow to another repository, update its `github.repository` guard.

You need at least one GitHub **Release** for `latest` URLs to resolve reliably.

Example shapes in packaged `version.json`:

- `announcement_url` → `https://github.com/LensDNS/LensDNS/releases/tag/<tag>`
- `download_*` → `https://github.com/LensDNS/LensDNS/releases/download/<tag>/<filename>`
