# SigmaDrive protocol plug-in artifacts

This public repository is the source for version-pinned Statically URLs:

`https://cdn.statically.io/gh/sigmadrive/SigmaDrive-Plugins@<tag>/<path>`

## Experimental S3 artifact

`poc/share-s3/0.1.0/windows-amd64/sigma-share-s3.exe` is an experimental Windows x64 build used to verify delivery through the real GitHub-to-Statically path. It is not a stable or generally supported product release. It was built from the private SigmaDrive workspace; the small, independently buildable public source package and license review are still pending.

The worker implements the read-only S3-compatible folder listing and file-read path. Its separate-origin local PoC passed an integrity check, TLS 1.3/SigV4 reads, a 950,272-byte file hash check, write rejection, and shutdown checks. This artifact's CDN download test only verifies delivery bytes; it does not make this file trusted by SigmaDrive clients.

SHA-256: `0261296b73d435951fa3fce26a45a3a85cbcc17646c542ac5cac0bd67438b1fa`

Do not use this experimental artifact for production data. Production installation requires the SigmaDrive client trust root and signed TUF metadata, which are not published in this proof-of-concept directory.
