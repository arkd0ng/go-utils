# random Package - Work Plan / 작업 계획

**Version**: v1.2.x Roadmap  
**Updated**: 2025-10-16

---

## Phase 1: Foundations (v1.2.000)
- [ ] Implement core random string helpers (alpha, alnum, hex)
- [ ] Implement numeric helpers (IntRange, FloatRange)
- [ ] Add deterministic seeding helper for tests
- [ ] Create basic unit tests covering length, alphabet, determinism
- [ ] Document usage in README and USER_MANUAL

## Phase 2: Deterministic Utilities (v1.2.001)
- [ ] Expose `NewRandWithSeed` helper for consumers
- [ ] Provide ability to capture/restore seeds
- [ ] Add examples demonstrating reproducible outputs
- [ ] Update docs with guidance for testing scenarios

## Phase 3: Secure Tokens (v1.2.002)
- [ ] Implement crypto-secure token generator (URL-safe)
- [ ] Provide helpers for API keys, short-lived codes
- [ ] Benchmark against standard library implementations
- [ ] Add security notes to docs

## Phase 4: Benchmarks & Polish (v1.2.003)
- [ ] Add Go benchmarks for string/number generators
- [ ] Optimise hotspots (allocation reduction)
- [ ] Finalise documentation and changelog entries
- [ ] Review package for public release readiness

---

> Track progress by checking items above as the package evolves.  
> 패키지 개발이 진행되면 위 항목을 체크하며 진행 상황을 추적하세요.
