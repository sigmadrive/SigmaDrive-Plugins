# SigmaDrive S3 share worker

`share-s3` is an optional, on-demand server worker. It is not a cloud-drive mount connector and is not bundled in the SigmaDrive installer.

## Artifact contract

- One standalone executable per plugin version and OS/CPU: `plugins/share-s3/<version>/<os>-<arch>/sigma-share-s3.exe`.
- Current experimental target: Windows x64, version `0.1.0`, worker ABI `1`.
- The executable is launched as a child process. SigmaDrive sends its bounded JSON configuration through stdin and reads the local TLS endpoint from stdout; credentials and folder paths are not command-line arguments.
- `-probe` prints the plugin ID, ABI, and version, then exits. This probe does not start a listener.
- The worker exposes a loopback-only, TLS 1.3 S3-compatible, read-only API. It starts only when an enabled S3 share needs it and stops on revoke, logout, or parent-process exit.
- Artifact bytes, platform, ABI, and SHA-256 are descriptive metadata. They do not authorize installation. The SigmaDrive client must verify the same target against its installed TUF trust root and threshold-signed metadata before staging or activation.
- Statically is the preferred artifact origin. On an artifact download failure, the client may retry the exact same immutable tag-relative path from `raw.githubusercontent.com`. Metadata and trust roots never fail over.

## Experimental status

This tag packages the previously tested Windows x64 PoC binary. It is not a production release. The minimal independently buildable public source, license review, SBOM, production TUF metadata/trust root, code signing, and end-to-end application installation configuration are not yet published. Do not use this artifact for production shares.
