# SigmaDrive protocol plug-in artifacts

This public repository is the source for version-pinned Statically URLs:

`https://cdn.statically.io/gh/sigmadrive/SigmaDrive-Plugins@<tag>/<path>`

## Module file contract

Each optional worker is versioned independently and published as a standalone OS/CPU-specific executable. The artifact contract, probe interface, runtime behavior, trust requirements, and current limitations are in [`plugins/share-s3/README.md`](plugins/share-s3/README.md). The current S3 target is `plugins/share-s3/0.1.0/windows-amd64/sigma-share-s3.exe`; its descriptor and checksum sit beside it. The TUF target mirror is under `targets/stable/windows-amd64/components/share-s3.exe`.

Statically is the primary binary origin. The SigmaDrive client may use the exact same tag-pinned GitHub raw file as a secondary origin if the CDN download fails. The GitHub copy is not a separate trust source: client installation still requires SigmaDrive's installed TUF trust root and signed target metadata.

## Experimental S3 artifact

`poc/share-s3/0.1.0/windows-amd64/sigma-share-s3.exe` is the original experimental Windows x64 delivery test. The versioned module path above is the same PoC build packaged using the agreed module contract; neither is a stable or generally supported product release. It was built from the private SigmaDrive workspace; the small, independently buildable public source package and license review are still pending.

The worker implements the read-only S3-compatible folder listing and file-read path. Its separate-origin local PoC passed an integrity check, TLS 1.3/SigV4 reads, a 950,272-byte file hash check, write rejection, and shutdown checks. This artifact's CDN download test only verifies delivery bytes; it does not make this file trusted by SigmaDrive clients.

SHA-256: `0261296b73d435951fa3fce26a45a3a85cbcc17646c542ac5cac0bd67438b1fa`

Do not use this experimental artifact for production data. Production installation requires the SigmaDrive client trust root and signed TUF metadata, which are not published in this proof-of-concept directory.
