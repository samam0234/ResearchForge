# AGENTS.md - ResearchForge Agent Guidelines (SSOT)

> **단일 진실 공급원 (Single Source of Truth, SSOT)**  
> ResearchForge 프로젝트에 참여하는 모든 AI 에이전트(Claude, Codex, Grok, Gemini, Antigravity, Cursor, Manus, Kimi, Ollama, OpenCode, Qwen, Windsurf 등)와 개발자는 본 문서의 규칙을 **예외 없이 엄격하게 준수**해야 합니다.

---

## 1. 프로젝트 개요 및 기술 스택

### 1.1 프로젝트 정보
- **프로젝트명**: ResearchForge
- **목표**: 사용자가 자연어로 연구 주제를 입력하면, 다중 AI 에이전트(Researcher, FactChecker, Analyst, Writer)가 협업하여 정보 수집, 출처 검증, 분석, 보고서 생성을 수행하는 멀티 에이전트 리서치 플랫폼.

### 1.2 시스템 아키텍처 및 서비스 구성
- **Frontend** (`frontend/`, Port: `3000`): Nuxt.js 3 (Vue 3, TypeScript, TailwindCSS)
- **Backend API** (`backend/`, Port: `8080`): Go (Fiber / Chi) - 인증, 작업 큐잉, 실시간 WebSocket 푸시
- **AI Server** (`ai-server/`, Port: `8000`): FastAPI (Python 3.10+) - CrewAI / LangGraph 기반 멀티 에이전트 오케스트레이션
- **Database**: PostgreSQL 16 (JSONB, pgvector 지원)
- **Cache & Message Queue**: Redis 7
- **Deployment**: Docker Compose (`researchforge`)

---

## 2. 필수 코딩 및 작업 수칙

### 2.1 커밋 메시지 규칙 (엄격 준수)
- **형식**: `type(scope): 한국어 요약`
- `type`과 `scope`는 영문 소문자 사용.
- **콜론 뒤 공백 1칸 필수**, 요약과 본문은 반드시 **한국어**로 작성.
- 허용되는 `type`:
  - `feat`: 새로운 기능 추가
  - `fix`: 버그 수정
  - `docs`: 문서 추가 및 수정
  - `style`: 코드 포맷팅, 세미콜론 누락 등 (코드 동작 변경 없음)
  - `refactor`: 리팩터링 (기능 추가나 버그 수정 없는 코드 변경)
  - `test`: 테스트 코드 추가 및 수정
  - `chore`: 빌드 업무 수정, 패키지 매니저 설정, 스캐폴딩 등
  - `perf`: 성능 향상을 위한 코드 변경

### 2.2 커밋 기록 문서화 (필수)
- 매 커밋 시마다 `Docs/branchs/commits/YYMMDD_HHMM_[id]_[이름]_[브랜치].md` 형식의 문서를 작성하여 **해당 커밋에 반드시 함께 포함**해야 합니다.
- 템플릿: `Docs/branchs/commits/_template.md` 사용.
- 커밋 기록 파일 누락 시 커밋하지 않습니다.

### 2.3 브랜치 관리 전략
- 모든 작업은 **오직 `feature/*` 브랜치**에서만 수행합니다.
- `main` 및 `develop` 브랜치는 직접 커밋을 금지하며 **병합(Merge) 전용**입니다.
- 병합 시에는 반드시 이력을 보존하기 위해 `--no-ff` 옵션을 사용합니다.
- **즉시 병합 금지**: 기능 작성이 끝났다고 임의로 `develop`나 `main`에 즉시 병합하지 않습니다. 사용자의 명시적인 지시가 있을 때만 최종 병합 작업을 일괄 수행합니다.

### 2.4 파괴적 작업 및 되돌리기 어려운 작업에 대한 사전 확인
- 다음 작업들은 **반드시 사용자의 명시적 지시나 승인**이 있을 때만 수행합니다:
  - 원격 저장소로의 `git push`
  - 브랜치 `merge` (병합)
  - 파일 또는 디렉토리의 영구 `delete` (삭제)
  - 기존 핵심 로직 또는 파일의 전면 덮어쓰기 (`overwrite`)
- 한 번 승인받은 내용을 다른 작업 맥락에 임의로 확장 적용하지 않습니다.

### 2.5 단일 진실 공급원(SSOT) 준수
- 에이전트별 설정 디렉토리(`.AI_Handover/.Claude/`, `.AI_Handover/.Cursor/`, `.AI_Handover/.Antigravity/` 등)에는 고유한 툴별 특화 설정만 기록하며, 공통 규칙은 반드시 본 `AGENTS.md`를 참조하도록 링크합니다.

---

## 3. 디렉토리 구조 및 역할

```
ResearchForge/
├── .AI_Handover/            # 각 AI 에이전트별 인수인계 가이드 (.Agent, .Claude, .Antigravity 등)
├── Docs/                    # 프로젝트 공식 문서 체계
│   ├── plan/                # 기획서, 인수인계서, DB 설계서 등
│   ├── Architecture/        # 아키텍처 다이어그램 및 설계 문서
│   ├── branchs/commits/     # 커밋별 기록 문서 (필수)
│   ├── find_debug/          # 디버깅 및 문제 해결 기록
│   ├── guidance/            # 실행 및 개발 환경 가이드
│   ├── trainings/           # 연구 및 학습 자료
│   ├── repeater/            # 인프라/네트워크 장애 및 복구 이력
│   ├── vaildates/           # 기능 검증 및 테스트 체크리스트
│   └── web_management/      # 웹 서비스 및 배포 관리 기록
├── frontend/                # Nuxt.js 3 프론트엔드
├── backend/                 # Go 백엔드 API 서버
├── ai-server/               # Python FastAPI 멀티 에이전트 서버
├── scripts/                 # 데이터베이스 초기화 및 자동화 스크립트
├── docker-compose.yml       # 전체 시스템 오케스트레이션
├── AGENTS.md                # [SSOT] 에이전트 작업 수칙 (본 파일)
├── CONTRIBUTING.md          # 기여 가이드라인
├── LICENSE                  # 라이선스 (MIT)
├── README.md                # 프로젝트 안내 문서
└── SECURITY.md              # 보안 정책
```

---

## 4. 에이전트 협업 및 로직 워크플로우

1. **Researcher Agent**:
   - 사용자의 쿼리를 바탕으로 다각도 웹 검색 수행 (Tavily / Serper / Scraper)
   - 주요 텍스트 데이터 및 URL 출처 수집
2. **FactChecker Agent**:
   - 수집된 원천 출처(Source URL) 검증 및 교차 검증
   - 신뢰할 수 없는 정보 배제 및 신뢰도 점수 부여
3. **Analyst Agent**:
   - 검증된 데이터를 바탕으로 핵심 인사이트, 비교 분석, 장단점 도출
4. **Writer Agent**:
   - 출처가 표기된 구조화된 Markdown 보고서 작성
   - 본문 내 인용 출처 하이라이트 및 참고문헌 목록 자동 구성

---

## 5. 변경 이력
- 2026-11-09: v1.0 초안 작성 (인수인계 및 기획서 v1.0 기준 확정)
