#!/usr/bin/env python3
"""Synchronize Markdown version strings with cfg/app.yaml.

The script looks for two kinds of targets:

1. Explicit placeholders `{Version}` placed inside Markdown files.
2. Existing version lines such as `**Version**: v1.11.000` or
   `**Version / 버전**: v1.11.000`.

Any file under the repository (recursively) with a `.md` suffix is processed,
so 새로운 문서를 추가해도 스크립트를 수정할 필요가 없습니다.
"""

from __future__ import annotations

import re
import sys
from pathlib import Path


CFG_PATH = Path("cfg/app.yaml")
VERSION_PLACEHOLDER = "{Version}"
VERSION_LINE_PATTERNS = [
    re.compile(r"(\*\*Version\s*/\s*버전\*\*:\s*)v\d+\.\d+\.\d+"),
    re.compile(r"(\*\*Version\*\*:\s*)v\d+\.\d+\.\d+"),
]
YAML_VERSION_REGEX = re.compile(r"^\s*version:\s*(v\d+\.\d+\.\d+)\s*$", re.MULTILINE)


def load_version() -> str:
    text = CFG_PATH.read_text(encoding="utf-8")
    match = YAML_VERSION_REGEX.search(text)
    if not match:
        raise SystemExit("Unable to locate app.version in cfg/app.yaml")
    return match.group(1)


def update_markdown(path: Path, version: str) -> bool:
    original = path.read_text(encoding="utf-8")
    updated = original

    if VERSION_PLACEHOLDER in updated:
        updated = updated.replace(VERSION_PLACEHOLDER, version)

    for pattern in VERSION_LINE_PATTERNS:
        updated = pattern.sub(rf"\1{version}", updated)

    if updated != original:
        path.write_text(updated, encoding="utf-8")
        return True
    return False


def main() -> int:
    version = load_version()
    markdown_files = [p for p in Path(".").rglob("*.md") if p.is_file()]

    any_changed = False
    for md_file in markdown_files:
        any_changed |= update_markdown(md_file, version)

    if any_changed:
        print(f"Updated version placeholders to {version}.")
    else:
        print("No Markdown files required updates.")

    return 0


if __name__ == "__main__":
    sys.exit(main())
