#!/usr/bin/env python3
"""Synchronize markdown version strings with cfg/app.yaml."""

import re
import sys
from pathlib import Path

import yaml


CFG_PATH = Path("cfg/app.yaml")


def load_version() -> str:
    data = yaml.safe_load(CFG_PATH.read_text(encoding="utf-8"))
    try:
        return data["app"]["version"]
    except KeyError as exc:
        raise SystemExit("cfg/app.yaml missing app.version") from exc


def update_file(path: Path, version: str) -> bool:
    text = path.read_text(encoding="utf-8")

    patterns = [
        re.compile(r"(\*\*Version\s*/\s*버전\*\*:\s*)v\d+\.\d+\.\d+"),
        re.compile(r"(\*\*Version\*\*:\s*)v\d+\.\d+\.\d+"),
    ]

    updated = text
    for pattern in patterns:
        updated = pattern.sub(rf"\1{version}", updated)

    if updated != text:
        path.write_text(updated, encoding="utf-8")
        return True
    return False


def main() -> int:
    version = load_version()

    target_files = [
        Path("README.md"),
        Path("docs/websvrutil/USER_MANUAL.md"),
        Path("docs/websvrutil/DEVELOPER_GUIDE.md"),
    ]

    changed = False
    for file_path in target_files:
        if not file_path.exists():
            continue
        changed |= update_file(file_path, version)

    if changed:
        print(f"Updated documentation version strings to {version}.")
    else:
        print("Documentation already up to date.")

    return 0


if __name__ == "__main__":
    sys.exit(main())
