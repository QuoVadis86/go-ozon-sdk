#!/usr/bin/env python3
"""Summarize changes between two Ozon Swagger JSON snapshots."""

import json
import sys
from pathlib import Path


HTTP_METHODS = {"get", "post", "put", "patch", "delete", "head", "options"}


def operations(spec):
    return {
        (path, method.upper()): operation
        for path, methods in spec["paths"].items()
        for method, operation in methods.items()
        if method.lower() in HTTP_METHODS
    }


def deprecated(operation):
    description = operation.get("description") or ""
    return bool(operation.get("deprecated")) or any(
        marker in description for marker in ("停用", "弃用", "废弃")
    )


def schema_ref(operation, source):
    if source == "request":
        body = operation.get("requestBody") or {}
    else:
        body = (operation.get("responses") or {}).get("200") or {}
    content = body.get("content") or {}
    return sorted(
        (media_type, (value.get("schema") or {}).get("$ref", ""))
        for media_type, value in content.items()
    )


def show(title, entries):
    print(f"\n{title} ({len(entries)}):")
    for key, detail in entries:
        print(f"  {key[1]} {key[0]}  {detail}")


def main(old_path, new_path):
    old = json.loads(Path(old_path).read_text())
    new = json.loads(Path(new_path).read_text())
    before, after = operations(old), operations(new)
    old_schemas = old["components"]["schemas"]
    new_schemas = new["components"]["schemas"]
    print(f"Operations: {len(before)} -> {len(after)}")
    print(f"Schemas: {len(old_schemas)} -> {len(new_schemas)}")

    added = sorted(after.keys() - before.keys())
    removed = sorted(before.keys() - after.keys())
    show("Added paths", [(key, after[key].get("operationId", "")) for key in added])
    show("Removed paths", [(key, before[key].get("operationId", "")) for key in removed])

    old_by_id = {
        op["operationId"]: key for key, op in before.items() if op.get("operationId")
    }
    moved = [
        (key, f"was {old_by_id[after[key]['operationId']]}")
        for key in added
        if after[key].get("operationId") in old_by_id
        and old_by_id[after[key]["operationId"]] != key
    ]
    show("Same operationId at a different path (check path corrections)", moved)

    newly_deprecated = [
        (key, after[key].get("operationId", ""))
        for key in sorted(after)
        if deprecated(after[key]) and (key not in before or not deprecated(before[key]))
    ]
    show("Newly deprecated", newly_deprecated)

    changed = []
    for key in sorted(after.keys() & before.keys()):
        differences = []
        for field in ("operationId", "tags"):
            if before[key].get(field) != after[key].get(field):
                differences.append(field)
        for source in ("request", "response"):
            if schema_ref(before[key], source) != schema_ref(after[key], source):
                differences.append(f"{source} schema ref")
        if differences:
            changed.append((key, ", ".join(differences)))
    show("Changed operation metadata or schema references", changed)

    print(f"\nSchema names added: {len(new_schemas.keys() - old_schemas.keys())}")
    print(f"Schema names removed: {len(old_schemas.keys() - new_schemas.keys())}")
    print(
        "Existing schemas changed: "
        + str(sum(old_schemas[name] != new_schemas[name] for name in old_schemas.keys() & new_schemas.keys()))
    )


if __name__ == "__main__":
    if len(sys.argv) != 3:
        sys.exit("usage: python3 scripts/swagger_diff.py OLD.json NEW.json")
    main(sys.argv[1], sys.argv[2])
