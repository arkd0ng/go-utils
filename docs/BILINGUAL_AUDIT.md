# Bilingual Audit Report / 이중 언어 감사 보고서

_Refreshed: 2025-10-17_

## Overview / 개요

- Purpose: identify files where English/Korean bilingual rules are not fully met or where explanatory comments need enrichment before the full documentation pass.  
  목적: 전체 문서 작업 전에 영문·한글 병기 규칙이 미흡하거나 주석 설명 보강이 필요한 파일을 식별합니다.
- Scope: all Go source files and Markdown documents under the repository.  
  범위: 저장소 내 모든 Go 소스 및 Markdown 문서.
- Method: static inspection scripts were run to detect comment lines containing only English or only Korean, followed by spot manual review of the most-affected files.  
  방법: 영문 또는 한글만 포함된 주석 라인을 찾는 스크립트를 실행하고, 영향도가 큰 파일을 수동 검토했습니다.

### Automated Metrics Snapshot / 자동 점검 지표

- Total Go files analysed: **258**  
  분석한 Go 파일 수: **258**
- Files with English-only comments: **231**  
  영문 전용 주석이 있는 파일 수: **231**
- Files with Korean-only comments: **185**  
  한글 전용 주석이 있는 파일 수: **185**
- Markdown files remain compliant with bilingual rules (no violations detected).  
  Markdown 문서는 현재 모두 병기 규칙을 준수하고 있습니다.

## High-Risk Files (EN-only comments ≥ 100) / 고위험 파일 (영문 전용 주석 100줄 이상)

| File / 파일 | EN-only | KO-only | Notes / 비고 |
| --- | ---:| ---:| --- |
| `websvrutil/middleware.go` | 304 | 79 | Recovery, logging, compression 관련 설명이 대부분 영어로만 존재 |
| `examples/mysql/main.go` | 193 | 95 | 예제 시나리오, 단계별 설명이 영어 위주 |
| `examples/maputil/main.go` | 190 | 149 | 로그 백업, 예제 단계 설명 대부분 단일 언어 |
| `random/string.go` | 175 | 148 | 문자셋 설명 및 함수 개요 영문/한글이 분리되어 있음 |
| `timeutil/timeutil_comprehensive_test.go` | 157 | 35 | 테스트 케이스 설명 다수가 영어만 제공 |
| `maputil/nested.go` | 154 | 55 | 함수 목적/동작 설명이 영어 전용 |
| `errorutil/error.go` | 153 | 83 | 패키지 개요와 함수 주석이 영어 중심 |
| `websvrutil/context_request.go` | 150 | 35 | 요청 정보/보안 섹션 영어 비율 높음 |
| `maputil/aggregate.go` | 147 | 34 | 알고리즘 설명, 복잡도 표기가 영어로만 제공 |
| `database/mysql/migration.go` | 136 | 30 | 마이그레이션 예제 및 경고 주석이 영어만 존재 |
| `database/mysql/schema.go` | 135 | 43 | 스키마 필드 설명 영어 중심 |
| `maputil/convert.go` | 132 | 34 | 변환 함수 설명 영어 위주 |
| `websvrutil/template.go` | 132 | 79 | 템플릿 엔진 구성요소 설명 영어 전용 다수 |
| `sliceutil/util.go` | 128 | 42 | 유틸 함수 주석이 한 언어로만 작성된 경우 많음 |
| `maputil/util.go` | 127 | 68 | 반복/탐색 관련 주석이 영어 위주 |
| `sliceutil/aggregate.go` | 127 | 19 | 집계 함수 요약이 영어만 제공 |
| `websvrutil/app.go` | 124 | 44 | App 구조 설명이 영어 전용 |
| `examples/httputil/main.go` | 121 | 85 | 예제 데이터 구조 설명 영어 중심 |
| `logging/logger.go` | 114 | 72 | 로거 사용 설명, 파라미터 문서화가 대부분 영어 |
| `logging/options.go` | 113 | 65 | 옵션 설명, 구성요소 소개가 영어 위주 |

> These files should be prioritised for manual bilingual rewriting where English-only comment blocks remain.  
> 위 파일들은 영어 전용 주석 블록이 많으므로 수동으로 병기 작업을 최우선으로 진행해야 합니다.

## High-Risk Files (KO-only comments ≥ 70) / 고위험 파일 (한글 전용 주석 70줄 이상)

| File / 파일 | KO-only | EN-only | Notes / 비고 |
| --- | ---:| ---:| --- |
| `examples/maputil/main.go` | 149 | 190 | 예제 진행 설명이 한쪽 언어로만 번갈아 존재 |
| `random/string.go` | 148 | 175 | 문자셋/보안 주석이 영어 또는 한글 한 줄씩만 존재 |
| `examples/mysql/main.go` | 95 | 193 | DB 시나리오 로그 안내가 언어별로 분리됨 |
| `examples/httputil/main.go` | 85 | 121 | REST 예제 주석이 단일 언어로 반복 |
| `errorutil/error.go` | 83 | 153 | 패키지 개요 및 함수 문서화가 영어 중심, 일부 한글만 있는 섹션도 존재 |
| `websvrutil/template.go` | 79 | 132 | 템플릿 엔진 구성 설명이 언어별로 불균형 |
| `websvrutil/middleware.go` | 79 | 304 | 미들웨어 설계 철학, 보안 설명이 영어 및 한글 단독으로 교차 |
| `examples/redis/main.go` | 73 | 98 | Redis 예제 단계 표현이 언어별로 나뉘어 있음 |
| `logging/logger.go` | 72 | 114 | 로깅 함수/옵션 설명이 충분히 병기되지 않음 |
| `maputil/util.go` | 68 | 127 | 반복/필터 유틸리티 주석이 영문 위주 |
| `logging/options.go` | 65 | 113 | 옵션 목록 설명이 거의 영어로만 제공 |

> Korean-only lines generally appear where an English counterpart already exists elsewhere in the file. When rewriting, mirror the richer explanation (often English) into Korean and keep both adjacent.  
> 한글 전용 주석은 주로 동일 파일의 다른 영어 설명과 분리되어 있으므로, 설명이 더 풍부한 쪽(대부분 영어)을 참고해 양쪽 언어를 인접하게 배치하는 것이 좋습니다.

## Medium Priority Files (20 ≤ EN-only < 100) / 중간 우선순위 (영문 전용 20~99줄)

- `fileutil/*` (특히 `path.go`, `info.go`, `dir.go`): 파일 시스템 유틸리티 설명이 영어 위주로 작성되어 있음.  
- `sliceutil/*`, `maputil/*`, `stringutil/*`: 알고리즘 복잡도, 예제, 경고 주석을 세부적으로 병기 필요.  
- `websvrutil` 테스트 및 헬퍼 (`context_*.go`, `session.go`, `validator.go` 등): 테스트 단계 안내와 보안 주석을 양 언어로 정비해야 함.  
- `httputil/*`: HTTP 요청 옵션, 응답 헬퍼의 파라미터 설명 병기가 부족.

각 파일은 동일한 패턴의 주석들이 반복되므로, 공통 템플릿을 정의하면 재작성 효율이 높아집니다.

## Additional Files Requiring Attention / 추가 검토가 필요한 파일

- `database/mysql` 패키지 전반: builder, client, config, export, metrics, pagination, retry 등 각 모듈에 영어 설명이 많이 남아 있습니다.  
  → 함수 목적, 파라미터 설명, 반환값 요약을 영어/한글 모두로 제공하도록 재작성 필요.
- `fileutil`, `maputil`, `sliceutil`, `stringutil`, `timeutil` 패키지: 알고리즘 복잡도, 예제, 경고 주석이 영어 위주로 작성되어 있어 세부 설명 보강 필요.  
  특히 `sliceutil/transform.go`, `stringutil/validation.go`는 영문만 존재하는 섹션이 다수입니다.
- `websvrutil` 테스트 코드 (`*_test.go`): 테스트 의도, 준비 단계, 검증 단계 설명이 영어 한 줄만으로 구성되어 있어 한글 설명 추가가 필요합니다.
- `examples/*`: 모든 예제 메인 파일에서 시나리오 안내, 단계별 로그 설명이 한 언어만 포함된 라인이 많습니다.  
  예) `examples/redis/main.go`의 “stringOperations demonstrates string operations” 등.
- `random/string.go`와 `random/string_test.go`: 문자셋 선택, 난수 품질 관련 설명을 영어/한글 모두로 풀어 써야 합니다.

## Markdown Documents / Markdown 문서

- 현재 `.md` 문서는 전체적으로 정합성이 양호하며, 영문/한글 병기 상태가 유지되고 있습니다.  
  다만 향후 내용 추가 시 “영문 단락 → 한글 단락” 또는 “문장 내 병기” 형식을 유지해야 합니다.

## Recommended Next Steps / 향후 작업 권장 사항

1. **Batch rewrite high-risk files**  
   최우선 파일(표 상단 20개)을 대상으로 영어 전용 주석을 찾아 문장 단위로 한글을 병기하고, 필요한 경우 설명을 더 친절하게 확장합니다.
2. **Establish comment templates**  
   패키지별로 반복되는 주석 패턴(예: “Example”, “Parameters”)에 대해 영어/한글 템플릿을 정의해 일괄 적용합니다.
3. **Re-run automated checks after edits**  
   수정 후 동일 스크립트를 다시 실행하여 영어 전용 라인이 0에 가깝도록 관리합니다.
4. **Address outstanding test failures**  
   `go test ./...` 실행 시 `httputil` 패키지의 버전 상수 불일치로 테스트가 실패하고 있으므로, 문서화 작업과 병행해 버전 기대값을 정리하는 것이 좋습니다.

## Appendix: Detection Method / 부록: 점검 방법

- Python 스크립트로 `//` 주석 라인을 탐색하여 영어/한글 포함 여부를 판별했습니다.  
- 동일 파일 내에서 영어 전용 주석이 발견된 경우 상위 3개 샘플 라인을 기록했습니다.  
- Markdown 문서는 영어 또는 한글이 결여된 경우에만 별도 보고하도록 검사했습니다.
- **Tests (`*_test.go`)**: 테스트 함수 설명이 대부분 영어 또는 한글 한 쪽에 치우쳐 있음. “준비 → 실행 → 검증” 흐름을 양 언어로 명확히 서술할 필요가 있습니다.
- **Examples (`examples/*`)**: 학습자가 직접 읽을 텍스트이므로, 로그/콘솔 출력 주석도 모두 병기되도록 보완해야 합니다.
- 개별 문서의 섹션 제목과 내용은 현재 이중 언어를 잘 유지하고 있으나, 향후 업데이트 시 “영문 단락 → 한글 단락” 또는 “문장 내 병기” 규칙을 준수해야 합니다.

## Recommendation Summary / 개선 권고 요약

1. **Target high-risk files first**  
   최우선 파일(영문 전용 ≥ 100 또는 한글 전용 ≥ 70)부터 주석을 재작성합니다.  
2. **Adopt bilingual comment templates**  
   “Parameters / 매개변수”, “Returns / 반환값”, “Example / 예제” 등 공통 절에 사용할 템플릿을 마련해 일관성을 확보합니다.  
3. **Enhance explanatory depth while translating**  
   단순 번역이 아닌, 설명이 부족한 경우(특히 `random`, `errorutil`, `websvrutil`)에는 배경/주의 사항까지 추가합니다.  
4. **Re-run automated checks**  
   수정 후 동일 스크립트를 재실행하여 영어 전용/한글 전용 라인이 남아 있는 파일을 지속적으로 추적합니다.  
5. **Resolve outstanding tests**  
   현재 `go test ./...` 실행 시 `httputil` 패키지에서 버전 불일치로 실패하므로, 문서/주석 개편과 병행해 버전 기대 값 정리를 권장합니다.

## Progress Log (2025-10-17) / 진행 현황

- `database/mysql/batch.go`: 모든 주석을 한 줄 내 영문/한글 병기로 통합하고 예제 코드에 보충 설명을 추가했습니다.  
  → 남은 영문 전용 주석 0줄.
- `websvrutil/middleware.go`: `Recovery` 섹션의 개요 주석을 양 언어로 재작성하고 예제 라인을 병기했으나, 나머지 미들웨어 주석은 후속 정리가 필요합니다.  
  → 차기 작업: Logger·CORS·RateLimiter 등 주요 섹션 주석 보강.
