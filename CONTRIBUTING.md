# Contributing to ResearchForge

ResearchForge 프로젝트에 기여해 주셔서 감사합니다. 원활한 협업과 품질 유지를 위해 아래 가이드라인을 반드시 숙지하고 준수해 주시기 바랍니다.

## 1. 단일 진실 공급원 (SSOT)

모든 작업 수칙의 기준은 루트 디렉토리의 `AGENTS.md`입니다. 개발 및 에이전트 작업 전 반드시 `AGENTS.md`를 확인하세요.

## 2. 브랜치 전략

- 모든 개발 작업은 오직 `feature/*` 브랜치에서만 진행합니다.
  - 예: `feature/scaffold-project`, `feature/agent-researcher`, `feature/report-viewer`
- `main`과 `develop` 브랜치는 직접 커밋할 수 없으며 병합(Merge) 전용입니다.
- 모든 병합은 히스토리 보존을 위해 `--no-ff` (No Fast-Forward) 방식으로 진행합니다.
- 임의 병합은 금지되며, 작업 단위가 완료된 후 사용자의 명시적 지시에 따라 총 병합을 수행합니다.

## 3. 커밋 메시지 컨벤션

커밋 메시지는 다음 규격을 엄격히 준수해야 합니다:

```
type(scope): 한국어 요약

[선택적 본문]
- 상세 작업 내용
```

- **허용 Type**:
  - `feat`: 새로운 기능 추가
  - `fix`: 버그 수정
  - `docs`: 문서 추가 및 변경
  - `style`: 코드 스타일 수정 (동작 변경 없음)
  - `refactor`: 리팩터링
  - `test`: 테스트 코드
  - `chore`: 빌드, 패키지, 환경설정 변경
  - `perf`: 성능 개선

## 4. 커밋 기록 문서화 (필수)

커밋을 생성할 때마다 해당 커밋의 작업 내역을 문서화하여 함께 커밋해야 합니다:
- **경로**: `Docs/branchs/commits/YYMMDD_HHMM_[id]_[이름]_[브랜치].md`
- **템플릿**: `Docs/branchs/commits/_template.md`

## 5. 파괴적 작업 방지

`push`, `merge`, 파일 `delete`, 덮어쓰기(`overwrite`) 등의 영향도가 큰 작업은 사용자의 명시적 승인을 거친 후에만 수행합니다.
