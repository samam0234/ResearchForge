# .AI_Handover/.Agent - 모든 에이전트 공통 가이드

본 디렉토리는 ResearchForge 프로젝트에 참여하는 모든 AI 에이전트(Claude, Cursor, Antigravity, Gemini 등)가 숙지해야 할 공통 원칙을 정의합니다.

## 단일 진실 공급원 (SSOT)
- 모든 규칙의 최상위 기준은 **프로젝트 루트의 `AGENTS.md`**입니다.
- 개별 에이전트 디렉토리는 본 파일 및 루트 `AGENTS.md`를 참조해야 합니다.

## 핵심 규칙 요약
1. **커밋 메시지**: `type(scope): 한국어 요약` 규격 엄수 (예: `feat(research): 검색 API 연동`)
2. **커밋 기록 문서**: 매 커밋 시 `Docs/branchs/commits/YYMMDD_HHMM_[id]_[이름]_[브랜치].md` 작성 필수
3. **브랜치 전략**: 작업은 오직 `feature/*` 브랜치에서만 수행 (`main`, `develop` 직접 커밋 금지)
4. **파괴적 작업 사전 승인**: push, merge, delete, 주요 파일 overwrite는 반드시 사용자 명시적 승인 후 실행
