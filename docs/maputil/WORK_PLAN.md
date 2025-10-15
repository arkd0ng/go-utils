# maputil Package Work Plan / maputil 패키지 작업 계획

## Project Information / 프로젝트 정보

- **Package Name / 패키지명**: maputil
- **Version / 버전**: v1.8.001
- **Start Date / 시작일**: 2025-10-15
- **Estimated Completion / 예상 완료**: 2025-10-15

## Overview / 개요

Implement a comprehensive map utilities package with 80+ type-safe functions across 10 categories, following the successful pattern of sliceutil.

sliceutil의 성공적인 패턴을 따라 10개 카테고리에 걸쳐 80개 이상의 타입 안전 함수를 가진 포괄적인 맵 유틸리티 패키지를 구현합니다.

## Work Phases / 작업 단계

### Phase 1: Setup and Core Infrastructure / 설정 및 핵심 인프라
**Estimated Time / 예상 시간**: 30 minutes

#### 1.1 Project Setup / 프로젝트 설정
- [x] Create maputil branch
- [x] Update version to v1.8.001 in cfg/app.yaml
- [x] Create directory structure
- [x] Write DESIGN_PLAN.md
- [x] Write WORK_PLAN.md

#### 1.2 Package Foundation / 패키지 기반
- [ ] Create maputil.go (package documentation, version, types)
- [ ] Define Entry[K, V] type
- [ ] Define Number and Ordered constraints
- [ ] Add package-level constants

### Phase 2: Core Functions Implementation / 핵심 함수 구현
**Estimated Time / 예상 시간**: 2 hours

#### 2.1 Basic Operations (11 functions)
**File: basic.go**

- [ ] `Get[K comparable, V any](m map[K]V, key K) (V, bool)`
- [ ] `GetOr[K comparable, V any](m map[K]V, key K, defaultValue V) V`
- [ ] `Set[K comparable, V any](m map[K]V, key K, value V) map[K]V`
- [ ] `Delete[K comparable, V any](m map[K]V, keys ...K) map[K]V`
- [ ] `Has[K comparable, V any](m map[K]V, key K) bool`
- [ ] `IsEmpty[K comparable, V any](m map[K]V) bool`
- [ ] `IsNotEmpty[K comparable, V any](m map[K]V) bool`
- [ ] `Len[K comparable, V any](m map[K]V) int`
- [ ] `Clear[K comparable, V any](m map[K]V) map[K]V`
- [ ] `Clone[K comparable, V any](m map[K]V) map[K]V`
- [ ] `Equal[K comparable, V comparable](m1, m2 map[K]V) bool`

#### 2.2 Transformation Functions (10 functions)
**File: transform.go**

- [ ] `Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R`
- [ ] `MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V`
- [ ] `MapValues[K comparable, V any, R any](m map[K]V, fn func(V) R) map[K]R`
- [ ] `MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, fn func(K1, V1) (K2, V2)) map[K2]V2`
- [ ] `Invert[K comparable, V comparable](m map[K]V) map[V]K`
- [ ] `Flatten[K comparable, V any](m map[K]map[string]V, delimiter string) map[string]V`
- [ ] `Unflatten[V any](m map[string]V, delimiter string) map[string]interface{}`
- [ ] `Chunk[K comparable, V any](m map[K]V, size int) []map[K]V`
- [ ] `Partition[K comparable, V any](m map[K]V, fn func(K, V) bool) (map[K]V, map[K]V)`
- [ ] `Compact[K comparable, V any](m map[K]V) map[K]V`

#### 2.3 Aggregation Functions (9 functions)
**File: aggregate.go**

- [ ] `Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R`
- [ ] `Sum[K comparable, V Number](m map[K]V) V`
- [ ] `Min[K comparable, V Ordered](m map[K]V) (K, V, bool)`
- [ ] `Max[K comparable, V Ordered](m map[K]V) (K, V, bool)`
- [ ] `MinBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)`
- [ ] `MaxBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool)`
- [ ] `Average[K comparable, V Number](m map[K]V) float64`
- [ ] `GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V`
- [ ] `CountBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G]int`

### Phase 3: Advanced Functions Implementation / 고급 함수 구현
**Estimated Time / 예상 시간**: 2 hours

#### 3.1 Merge Operations (8 functions)
**File: merge.go**

- [ ] `Merge[K comparable, V any](maps ...map[K]V) map[K]V`
- [ ] `MergeWith[K comparable, V any](fn func(V, V) V, maps ...map[K]V) map[K]V`
- [ ] `DeepMerge(maps ...map[string]interface{}) map[string]interface{}`
- [ ] `Union[K comparable, V any](maps ...map[K]V) map[K]V`
- [ ] `Intersection[K comparable, V any](maps ...map[K]V) map[K]V`
- [ ] `Difference[K comparable, V any](m1, m2 map[K]V) map[K]V`
- [ ] `SymmetricDifference[K comparable, V any](m1, m2 map[K]V) map[K]V`
- [ ] `Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V`

#### 3.2 Filter Operations (7 functions)
**File: filter.go**

- [ ] `Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V`
- [ ] `FilterKeys[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V`
- [ ] `FilterValues[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V`
- [ ] `Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V`
- [ ] `Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V`
- [ ] `OmitBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V`
- [ ] `PickBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V`

#### 3.3 Conversion Functions (8 functions)
**File: convert.go**

- [ ] `Keys[K comparable, V any](m map[K]V) []K`
- [ ] `Values[K comparable, V any](m map[K]V) []V`
- [ ] `Entries[K comparable, V any](m map[K]V) []Entry[K, V]`
- [ ] `FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V`
- [ ] `FromSlice[K comparable, V any](slice []V, fn func(V) K) map[K]V`
- [ ] `FromSliceBy[K comparable, V any, R any](slice []V, keyFn func(V) K, valueFn func(V) R) map[K]R`
- [ ] `ToSlice[K comparable, V any, R any](m map[K]V, fn func(K, V) R) []R`
- [ ] `ToJSON[K comparable, V any](m map[K]V) (string, error)`

### Phase 4: Specialized Functions Implementation / 특수 함수 구현
**Estimated Time / 예상 시간**: 1.5 hours

#### 4.1 Predicate Checks (7 functions)
**File: predicate.go**

- [ ] `Every[K comparable, V any](m map[K]V, fn func(K, V) bool) bool`
- [ ] `Some[K comparable, V any](m map[K]V, fn func(K, V) bool) bool`
- [ ] `None[K comparable, V any](m map[K]V, fn func(K, V) bool) bool`
- [ ] `HasKey[K comparable, V any](m map[K]V, key K) bool`
- [ ] `HasValue[K comparable, V comparable](m map[K]V, value V) bool`
- [ ] `HasEntry[K comparable, V comparable](m map[K]V, key K, value V) bool`
- [ ] `IsSubset[K comparable, V comparable](subset, superset map[K]V) bool`

#### 4.2 Key Operations (8 functions)
**File: keys.go**

- [ ] `KeysSlice[K comparable, V any](m map[K]V) []K`
- [ ] `KeysSorted[K Ordered, V any](m map[K]V) []K`
- [ ] `KeysBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []K`
- [ ] `RenameKey[K comparable, V any](m map[K]V, oldKey, newKey K) map[K]V`
- [ ] `RenameKeys[K comparable, V any](m map[K]V, mapping map[K]K) map[K]V`
- [ ] `SwapKeys[K comparable, V any](m map[K]V, key1, key2 K) map[K]V`
- [ ] `FindKey[K comparable, V any](m map[K]V, fn func(K, V) bool) (K, bool)`
- [ ] `FindKeys[K comparable, V any](m map[K]V, fn func(K, V) bool) []K`

#### 4.3 Value Operations (7 functions)
**File: values.go**

- [ ] `ValuesSlice[K comparable, V any](m map[K]V) []V`
- [ ] `ValuesSorted[K comparable, V Ordered](m map[K]V) []V`
- [ ] `ValuesBy[K comparable, V any](m map[K]V, fn func(K, V) bool) []V`
- [ ] `UniqueValues[K comparable, V comparable](m map[K]V) []V`
- [ ] `FindValue[K comparable, V any](m map[K]V, fn func(K, V) bool) (V, bool)`
- [ ] `ReplaceValue[K comparable, V comparable](m map[K]V, oldValue, newValue V) map[K]V`
- [ ] `UpdateValues[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V`

#### 4.4 Comparison Functions (6 functions)
**File: comparison.go**

- [ ] `EqualFunc[K comparable, V any](m1, m2 map[K]V, eq func(V, V) bool) bool`
- [ ] `Diff[K comparable, V comparable](m1, m2 map[K]V) map[K]V`
- [ ] `DiffKeys[K comparable, V any](m1, m2 map[K]V) []K`
- [ ] `CommonKeys[K comparable, V any](maps ...map[K]V) []K`
- [ ] `AllKeys[K comparable, V any](maps ...map[K]V) []K`
- [ ] `Compare[K comparable, V comparable](m1, m2 map[K]V) (added, removed, modified map[K]V)`

### Phase 5: Testing / 테스트
**Estimated Time / 예상 시간**: 3 hours

#### 5.1 Unit Tests
- [ ] basic_test.go (11 functions)
- [ ] transform_test.go (10 functions)
- [ ] aggregate_test.go (9 functions)
- [ ] merge_test.go (8 functions)
- [ ] filter_test.go (7 functions)
- [ ] convert_test.go (8 functions)
- [ ] predicate_test.go (7 functions)
- [ ] keys_test.go (8 functions)
- [ ] values_test.go (7 functions)
- [ ] comparison_test.go (6 functions)

#### 5.2 Benchmark Tests
- [ ] Benchmark all performance-critical functions
- [ ] Memory allocation tests
- [ ] Comparison with standard library approaches

#### 5.3 Coverage
- [ ] Achieve 100% code coverage
- [ ] Test edge cases (empty maps, nil values, single element)
- [ ] Test type variations (string, int, float64, custom types)

### Phase 6: Examples / 예제
**Estimated Time / 예상 시간**: 1 hour

#### 6.1 Create examples/maputil/main.go
- [ ] Basic operations demonstration
- [ ] Transformation examples
- [ ] Aggregation examples
- [ ] Merge operations examples
- [ ] Filter operations examples
- [ ] Conversion examples
- [ ] Predicate checks examples
- [ ] Key/Value operations examples
- [ ] Comparison examples
- [ ] Real-world use cases

### Phase 7: Documentation / 문서화
**Estimated Time / 예상 시간**: 4 hours

#### 7.1 Package Documentation
- [ ] maputil/README.md
  - Quick start guide
  - Function categorization
  - Common use cases
  - Performance notes

#### 7.2 Comprehensive Manuals
- [ ] docs/maputil/USER_MANUAL.md (~3000+ lines)
  - Table of Contents
  - Introduction
  - Installation
  - Quick Start
  - Function Reference (all 81 functions)
  - Usage Patterns
  - Common Use Cases
  - Best Practices
  - Troubleshooting
  - FAQ

- [ ] docs/maputil/DEVELOPER_GUIDE.md (~2000+ lines)
  - Architecture Overview
  - Package Structure
  - Core Components
  - Internal Implementation
  - Design Patterns
  - Adding New Features
  - Testing Guide
  - Performance
  - Contributing Guidelines
  - Code Style

- [ ] docs/maputil/PERFORMANCE_BENCHMARKS.md
  - Benchmark results
  - Performance comparison
  - Memory usage
  - Optimization tips

### Phase 8: Integration and Finalization / 통합 및 마무리
**Estimated Time / 예상 시간**: 1 hour

#### 8.1 Update Root Documentation
- [ ] Update README.md
  - Add maputil section
  - Update version history
  - Add usage examples

#### 8.2 Update CHANGELOG
- [ ] Create docs/CHANGELOG/CHANGELOG-v1.8.md
- [ ] Update CHANGELOG.md with v1.8.x overview

#### 8.3 Update CLAUDE.md
- [ ] Add maputil architecture section
- [ ] Add development guidelines for maputil

#### 8.4 Final Checks
- [ ] Run all tests: `go test ./maputil -v`
- [ ] Check coverage: `go test ./maputil -cover`
- [ ] Run benchmarks: `go test ./maputil -bench=.`
- [ ] Build examples: `go build ./examples/maputil`
- [ ] Verify all documentation links
- [ ] Code review checklist

### Phase 9: Git Operations / Git 작업
**Estimated Time / 예상 시간**: 30 minutes

#### 9.1 Commit and Push
- [ ] Stage all changes
- [ ] Commit with message: "Feat: Add comprehensive maputil package v1.8.001"
- [ ] Push to maputil branch
- [ ] Verify on GitHub

#### 9.2 Merge to Main (Optional)
- [ ] Create pull request
- [ ] Review changes
- [ ] Merge to main branch
- [ ] Tag release v1.8.001

## Checklist Summary / 체크리스트 요약

### Code Implementation / 코드 구현
- [ ] 10 source files (maputil.go + 9 category files)
- [ ] 81 functions total
- [ ] All functions with bilingual documentation
- [ ] Type-safe with generics
- [ ] Immutable operations

### Testing / 테스트
- [ ] 10 test files
- [ ] 100% code coverage
- [ ] Edge cases tested
- [ ] Benchmark tests
- [ ] Type variation tests

### Documentation / 문서화
- [ ] DESIGN_PLAN.md
- [ ] WORK_PLAN.md
- [ ] README.md (package)
- [ ] USER_MANUAL.md (~3000+ lines)
- [ ] DEVELOPER_GUIDE.md (~2000+ lines)
- [ ] PERFORMANCE_BENCHMARKS.md
- [ ] Updated root README.md
- [ ] Updated CHANGELOG

### Examples / 예제
- [ ] examples/maputil/main.go
- [ ] All 10 categories demonstrated
- [ ] Real-world use cases
- [ ] Bilingual comments

## Time Estimation / 시간 예상

| Phase | Description | Time |
|-------|-------------|------|
| 1 | Setup and Core Infrastructure | 30m |
| 2 | Core Functions (30 functions) | 2h |
| 3 | Advanced Functions (23 functions) | 2h |
| 4 | Specialized Functions (28 functions) | 1.5h |
| 5 | Testing | 3h |
| 6 | Examples | 1h |
| 7 | Documentation | 4h |
| 8 | Integration and Finalization | 1h |
| 9 | Git Operations | 30m |
| **Total** | | **~15.5 hours** |

## Success Criteria / 성공 기준

1. **Functionality / 기능성**
   - ✅ All 81 functions implemented
   - ✅ Type-safe with generics
   - ✅ Immutable operations
   - ✅ Zero dependencies

2. **Quality / 품질**
   - ✅ 100% test coverage
   - ✅ All benchmarks pass
   - ✅ No lint errors
   - ✅ Clean code

3. **Documentation / 문서화**
   - ✅ Comprehensive manuals
   - ✅ Bilingual documentation
   - ✅ Clear examples
   - ✅ Performance notes

4. **Usability / 사용성**
   - ✅ Intuitive API
   - ✅ 20+ lines → 1-2 lines reduction
   - ✅ Clear error messages
   - ✅ Easy to integrate

## Notes / 참고사항

- Follow sliceutil's successful pattern
- sliceutil의 성공적인 패턴 따르기
- Maintain consistency with existing packages
- 기존 패키지와 일관성 유지
- Prioritize performance and type safety
- 성능과 타입 안전성 우선
- Comprehensive testing is critical
- 포괄적인 테스트가 중요

## Next Steps / 다음 단계

After completion of maputil v1.8.x:
1. Consider fileutil package (v1.9.x)
2. Consider httputil package (v1.10.x)
3. Continue expanding go-utils ecosystem

maputil v1.8.x 완료 후:
1. fileutil 패키지 고려 (v1.9.x)
2. httputil 패키지 고려 (v1.10.x)
3. go-utils 생태계 지속 확장
