# SDK maintenance instructions

This repository is a Go 1.18 Ozon Seller SDK. Read [docs/SDK_MAINTENANCE.md](docs/SDK_MAINTENANCE.md) **before** replacing `origin/swagger.json` or running `go run gen.go`. That guide is the handoff for a new maintainer or coding agent given only a new Swagger file.

## Where to work

| Concern | Files |
| --- | --- |
| Source snapshot and diff report | `origin/swagger.json`, `scripts/swagger_diff.py` |
| Generated methods, models, integration tests | `*/service.go`, `*/types.go`, `*/service_test.go` |
| Generator and ordering rules | `gen.go`, `gen_order_test.go` |
| Product attributes V4 response variants | `product/attributes_v4_json.go`, `product/attributes_v4_test.go` |
| Retired public API compatibility | `product/attributes_v3_compat.go`, `returns/legacy_rfbs.go` |
| Deprecation detection | `deprecation_test.go` |
| Language header | `transport/language.go`, `transport/client.go`, `client.go` |

## Invariants for every Swagger refresh

- Treat Swagger as input, not unquestioned runtime truth. The V4 product attributes schema still contradicts observed JSON: `images` may contain URL strings or objects; `sku` and `total` may be JSON numbers or strings. Preserve the custom decoders and the generated image `URL` field.
- Preserve the public `NewClient` signature. The language header uses typed constants and `WithLanguage(ctx, ...)` for per-request selection; never require arbitrary header strings or lock all requests to one language at initialization.
- Compare old and new paths, operation IDs, tags, request/response references, and deprecation flags **before** generating. A missing path is not proof that the live API stopped working. A changed tag may move a method to another SDK package.
- Do not silently remove an existing public method or type. Retain retired methods in a non-generated compatibility file where feasible, with a standalone `Deprecated:` Go doc paragraph. State clearly when runtime availability is unverified. If a method moves packages, audit callers of its old package too.
- Known gap at v0.3.2: five exemplar operations moved from `beta.Service` to `fbs.Service` after a Swagger tag change; old `beta.Service` wrappers have not yet been restored. See the maintenance guide before another release.
- Keep explicit generator exceptions for real upstream mismatches, including stable schema names, the `TopSku` path, and the product image `URL` field. Check them against new evidence before changing them.
- The generator preserves existing method/type/field order and sorts additions. Re-running with the same Swagger must leave generated files byte-identical.
- Keep credentials, cookies, and signed documentation URLs out of the repository and tool output. Commit only the public Swagger JSON.

## Required verification

Run `python3 scripts/swagger_diff.py OLD.json origin/swagger.json` and review every removed or relocated operation. Then run `go run gen.go`, `go test gen.go gen_order_test.go`, and `env -u OZON_CLIENT_ID -u OZON_API_KEY go test ./...`. Run the generator a second time and confirm no generated diff changes. Review `git diff --check`, the public method/type diff, deprecation notices, the V4 attributes regression tests, and both README coverage tables before release.

The integration tests in generated `service_test.go` can call real Ozon endpoints when credentials are set. Use the credential-free command above for routine maintenance. Release tags point to the merged commit, not the feature branch.
