# .AI_Handover/.Claude - Claude 전용 인수인계 및 실행 가이드

## 1. 규칙의 단일 출처 (SSOT)
본 환경에서 작업하는 모든 에이전트는 루트 디렉토리의 **AGENTS.md** 및 .AI_Handover/.Agent/README.md를 최우선 기준으로 준수합니다.

## 2. Claude 작업 시 주의사항
- **커밋 메시지 규칙 준수**: `type(scope): 한국어 요약` 형식 엄수.
- **커밋 기록 문서 필수**: 커밋 전 반드시 `Docs/branchs/commits/` 폴더에 커밋 기록 파일 생성.
- **브랜치 규칙**: `feature/*` 브랜치에서만 작업하고 `develop`/`main`에 임의 병합 금지.
- **사용자 승인 필수**: push, merge, delete 작업은 사용자 승인 전 절대 수행 금지.
