# ResearchForge 인수인계 문서

**AI 기반 멀티 에이전트 리서치 도우미**

- 문서 버전: v1.0
- 작성일: 2026-11-09
- 프로젝트 상태: 기획 단계 → 개발 착수 준비
- SSOT: 루트 `AGENTS.md`
- 관련 문서: `Docs/plan/ResearchForge_기획서_v1.0.md`

---

## 1. 프로젝트 개요

### 1.1 프로젝트명
**ResearchForge**

### 1.2 한 줄 소개
사용자가 연구 주제나 질문만 입력하면, 여러 AI 에이전트가 협업하여 정보 수집 · 검증 · 분석 · 보고서 작성을 자동으로 수행하는 멀티 에이전트 리서치 플랫폼.

### 1.3 기술 스택 (확정)

| 구분 | 기술 | 비고 |
|------|------|------|
| Frontend | Nuxt.js 3 (Vue 3) | SSR/SSG, 파일 기반 라우팅 |
| Backend | Go (Fiber 또는 Chi) | 고성능, 동시성, 작업 큐잉 |
| AI Server | FastAPI (Python) | 에이전트 실행, 도구 호출 |
| AI Layer | CrewAI (MVP) → LangGraph (고도화) | 멀티 에이전트 오케스트레이션 |
| Database | PostgreSQL | JSONB, 향후 pgvector |
| Cache / Queue | Redis | 세션, 작업 상태, 실시간 진행률 |
| Search Tool | Tavily / Serper / DuckDuckGo | 웹 검색 |
| Auth | JWT + OAuth2 | |
| Deployment | Docker Compose (MVP) → K8s | |

### 1.4 핵심 가치
- **신뢰성**: 출처 기반 검증 + 다중 에이전트 교차 검증
- **효율성**: 주제 입력 → 구조화된 보고서 자동화
- **확장성**: 에이전트 역할·도구 추가 용이
- **투명성**: 각 에이전트 작업 과정과 근거 노출

---

## 2. 코딩 수칙 (필수 · 예외 없음)

### 2.1 커밋 메시지 규칙
- **형식**: `type(scope): 한국어 요약`
- type/scope는 영어, 제목 요약과 본문은 한국어
- 허용 type: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`, `perf`

**올바른 예시**
```
feat(research): 리서치 요청 API 및 작업 큐잉 구현
fix(agent): FactChecker 출처 검증 실패 시 재시도 로직 추가
docs(handover): ResearchForge 인수인계 문서 초안 작성
```

**잘못된 예시 (❌)**
```
update stuff
feat: add research feature (영어 요약)
수정함
```

### 2.2 커밋 기록 문서화 (필수)
- 커밋마다 `Docs/branchs/commits/YYMMDD_HHMM_[id]_[이름]_[브랜치].md` 파일을 남긴다.
- 템플릿은 `Docs/branchs/commits/_template.md`를 사용한다.
- 기록 파일도 동일 브랜치에 함께 커밋한다.

### 2.3 브랜치 전략
- 작업은 **오직 `feature/*` 브랜치**에서만 수행한다.
- `develop` / `main`은 병합 전용.
- 병합은 `--no-ff`로 이력을 보존한다.
- **한 브랜치가 끝날 때마다 즉시 병합하지 않는다.** 요청받은 작업을 모두 커밋한 뒤, 사용자 명시적 지시가 있을 때만 한 번에 총 병합한다.
- 병합 기록 문서도 총 병합 시점에 한 번에 작성한다.

### 2.4 되돌리기 어려운 작업 사전 확인
- `push`, 병합, 삭제, 덮어쓰기는 **사용자 명시적 지시가 있을 때만** 수행한다.
- 한 번 승인한 내용을 다른 상황에 임의로 확장 적용하지 않는다.

### 2.5 규칙의 단일 출처 (SSOT)
- 루트 `AGENTS.md`에 공유 규칙을 둔다.
- 에이전트별 파일(`.claude/CLAUDE.md`, `.cursor/rules` 등)은 `AGENTS.md`를 가리키기만 한다.
- 규칙 중복으로 인한 불일치를 방지한다.

### 2.6 스택·실행 환경 요약
- Frontend: `frontend/` (Nuxt 3, 포트 3000)
- Backend: `backend/` (Go, 포트 8080)
- AI Server: `ai-server/` (FastAPI, 포트 8000)
- Docker Compose 프로젝트명: `researchforge`
- 의존성: `package.json`, `go.mod`, `requirements.txt` / `pyproject.toml`

---

## 3. Docs 폴더 구조 (필수)

```
Docs/
├── plan/                  # 초기 기획·로직·DB·배포 계획
├── Architecture/          # 시스템·계층·데이터 흐름 아키텍처
├── branchs/
│   └── commits/           # 커밋 기록 (필수 하위 폴더)
├── find_debug/            # 작업/실행 디버깅 일일 기록
├── guidance/              # 기능·실행 가이드
├── trainings/             # 학습·연구 정리
├── repeater/              # 서버 연결·인프라 장애와 해결
├── vaildates/             # 검증·체크리스트
└── web_management/        # 배포 전/배포 시 웹 관리 기록
```

---

## 4. AI_Handover 에이전트 폴더 구조 (필수)

```
.AI_Handover/
├── .Agent/                # 모든 에이전트 공통
├── .Claude/               # Claude 전용
├── .Codex/                # Codex 전용
├── .Grok/                 # Grok 전용
├── .Gemini/               # Gemini 전용
├── .Antigravity/          # Antigravity 전용
├── .Cursor/               # Cursor 전용
├── .Manus/                # Manus 전용
├── .Kimi/                 # Kimi 전용
├── .Ollama/               # Ollama 전용
├── .Opencode/             # OpenCode 전용
├── .Qwen/                 # Qwen 전용
└── .Windsurf/             # Windsurf 전용
```

각 에이전트 폴더에는 최소한 다음을 포함한다:
- `README.md` (해당 에이전트 사용 시 주의사항)
- 공통 규칙은 `.Agent` 또는 루트 `AGENTS.md`를 참조하도록 작성

---

## 5. 시스템 아키텍처 및 로직 구조

### 5.1 전체 구조

```
[사용자]
    ↓
[Nuxt.js Frontend]  ←→  [Go Backend API]
                              ↓
                        [PostgreSQL + Redis]
                              ↓
                        [FastAPI AI Server]
                              ↓
                    [CrewAI / LangGraph]
                         /    |    \
                  [Search] [Verify] [Writer]
```

### 5.2 에이전트 워크플로우 (MVP - CrewAI)

```
Researcher → FactChecker → Analyst → Writer → (Critic 선택)
```

| 에이전트 | 역할 | 주요 도구 |
|----------|------|-----------|
| Researcher | 웹 검색 및 정보 수집 | Search Tool, Web Scraper |
| FactChecker | 출처 검증 및 교차 확인 | Search, Citation Extractor |
| Analyst | 정보 분석 및 인사이트 도출 | LLM Reasoning |
| Writer | 구조화된 보고서 작성 | Template Engine |
| Critic | 최종 검토 및 품질 개선 | Self-reflection |

- 순차 실행을 기본으로 하되, Researcher는 병렬 검색 가능
- 각 에이전트는 명확한 `role`, `goal`, `backstory` 정의
- 출력은 구조화된 JSON 또는 Markdown으로 강제
- 출처는 반드시 URL + 인용문 포함

### 5.3 데이터 모델 핵심

- `users`
- `research_projects` (status: pending/running/completed/failed)
- `agent_executions`
- `reports` (content_markdown, sources JSONB)

### 5.4 통신
- Frontend ↔ Backend: REST + WebSocket (실시간 진행 상태)
- Backend ↔ AI Server: 내부 REST (향후 gRPC 고려)

---

## 6. 필수 생성 파일

다음 파일은 프로젝트 루트에 반드시 생성한다:

- `.gitattributes`
- `LICENSE` (MIT 권장)
- `SECURITY.md`
- `CONTRIBUTING.md`
- `AGENTS.md` (SSOT)
- `README.md`
- `docker-compose.yml`

---

## 7. 개발 단계 (참고)

| 단계 | 기간 | 목표 |
|------|------|------|
| Phase 0 | 1주 | 환경 구축 및 프로토타입 |
| Phase 1 | 2주 | MVP 핵심 기능 (리서치 → 보고서 E2E) |
| Phase 2 | 2주 | 실시간 UI + 히스토리 |
| Phase 3 | 1~2주 | 품질 개선 및 배포 |
| Phase 4 | 이후 | LangGraph 전환, RAG, 팀 기능 |

### MVP 완료 기준
- [ ] 주제 입력 후 보고서 생성 성공
- [ ] 최소 3개 에이전트 협업 확인
- [ ] 출처 5개 이상 포함
- [ ] 프론트에서 결과 확인 가능

---

## 8. 에이전트 작업 시 주의사항

1. 모든 작업은 `feature/*` 브랜치에서 수행한다.
2. 커밋 시 반드시 커밋 기록 문서를 `Docs/branchs/commits/`에 함께 남긴다.
3. `push`, 병합, 삭제 작업은 사용자 명시적 지시 없이 수행하지 않는다.
4. 규칙 관련 질문은 항상 루트 `AGENTS.md`를 우선 참조한다.
5. 에이전트별 특수 설정은 해당 `.Claude`, `.Cursor` 등의 폴더에만 작성하고, 공통 규칙은 중복하지 않는다.

---

## 9. 참고 및 학습 포인트

- Nuxt 3 Composition API + Server Routes
- Go Fiber/Chi + 구조화된 에러 처리
- FastAPI + Pydantic v2
- CrewAI 에이전트/태스크/크루 설계
- LangGraph StateGraph와 조건부 엣지
- WebSocket 실시간 통신
- PostgreSQL JSONB + pgvector

---

**문서 끝**

이 문서는 ResearchForge 프로젝트의 단일 인수인계 문서로, 모든 에이전트와 개발자가 준수해야 할 수칙과 구조를 정의한다. 실제 구현 과정에서 세부 사항은 조정될 수 있으나, 커밋·브랜치·SSOT 규칙은 예외 없이 적용한다.
