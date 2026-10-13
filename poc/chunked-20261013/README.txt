SigmaDrive public chunk distribution experiment, 2026-10-13.
NOT a production plugin, embedding model or installation authorization.
32 MiB deterministic synthetic model; 4 MiB SHA-256 addressed chunks.
0.2.0 baseline; 0.2.1 changes model chunks 1 and 5; 0.2.2 changes 2 and 6.
0.2.3 deliberately fails the health probe to test preservation of 0.2.2.
The native fixture only reads config.json and model.bin and prints health JSON.
Root/targets use two ephemeral Ed25519 signatures. Other role keys are separate.
Public metadata expires seven days after preparation; private keys are never saved.
Use an out-of-band reviewed root and a full commit-pinned artifact URL.
No production app source, user data, credentials or production signing keys are included.
The embedded Go runtime license is in fixture-source/GO-LICENSE.txt.
