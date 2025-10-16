# random Package - Design Plan / 설계 계획

**Version**: v1.2.000 (initial draft)  
**Created**: 2025-10-16  
**Package**: `github.com/arkd0ng/go-utils/random`

---

## Overview / 개요

- Helper utilities for generating random strings, numbers, tokens with sensible defaults.  
  개발자가 자주 사용하는 랜덤 문자열, 숫자, 토큰 등을 손쉽게 생성하도록 돕습니다.
- Focus on simplicity, deterministic seeding for tests, and cryptographically secure helpers where necessary.  
  단순성과 테스트를 위한 결정적 시드, 필요한 경우 암호학적 안전성을 중점으로 합니다.

## Goals / 목표

1. Provide convenience wrappers on top of `crypto/rand` and `math/rand`.  
   `crypto/rand`, `math/rand`에 대한 편의 래퍼를 제공합니다.
2. Offer reusable presets (URL-safe tokens, human-readable codes).  
   URL-safe 토큰, 사람이 읽기 쉬운 코드 등 재사용 가능한 프리셋을 제공합니다.
3. Keep API deterministic-friendly (seed control) for tests.  
   테스트를 위해 시드를 제어할 수 있도록 합니다.
4. Ensure zero external dependencies.  
   외부 의존성이 없도록 합니다.

## Scope / 범위

- Random string/token generators with selectable alphabets.  
  알파벳을 선택할 수 있는 랜덤 문자열/토큰 생성기.
- Numeric helpers (int range, float range).  
  숫자 범위 헬퍼 (정수/실수).
- Utility to seed and restore deterministic state.  
  결정적 상태를 시드하고 복원하는 유틸리티.

## Out of Scope / 범위 외

- Full cryptographic library implementation.  
  암호화 라이브러리 전체 구현.
- Password hashing/salting (belongs in security-focused package).  
  패스워드 해싱/솔팅 기능.

## Roadmap / 로드맵 (Draft)

1. v1.2.000 – Initial helpers (strings, numbers) + docs.  
   초기 헬퍼(문자열, 숫자) 및 문서.
2. v1.2.001 – Add deterministic seeding utilities.  
   결정적 시드 유틸리티 추가.
3. v1.2.002 – Provide crypto-secure token helpers.  
   암호학적으로 안전한 토큰 헬퍼 제공.
4. v1.2.003 – Benchmark and optimise hot paths.  
   사용 빈도가 높은 경로 벤치마크 및 최적화.

> This document will be refined as the package evolves.  
> 패키지 개발 진행에 따라 문서를 지속적으로 업데이트합니다.
