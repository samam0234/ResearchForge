# ResearchForge 🔬

> **AI 기반 멀티 에이전트 리서치 도우미 플랫폼**  
> 사용자가 연구 주제나 질문을 입력하면, 여러 전문 AI 에이전트가 협업하여 정보 수집, 출처 검증, 분석, 구조화된 보고서 작성을 자동으로 수행합니다.

---

## 🌟 주요 특징

- **멀티 에이전트 파이프라인**: Researcher(검색) → FactChecker(검증) → Analyst(분석) → Writer(보고서 작성)의 분업화된 워크플로우
- **신뢰할 수 있는 출처 기반 리서치**: 단순 생성이 아닌 실제 웹 출처 및 인용문 교차 검증
- **모던 풀스택 아키텍처**:
  - **Frontend**: Nuxt.js 3 (Vue 3, TypeScript, TailwindCSS)
  - **Backend**: Go Fiber (고성능 비동기 API, 작업 큐, 실시간 WebSocket 푸시)
  - **AI Server**: FastAPI + Python (CrewAI / LangGraph 오케스트레이션)
  - **Database & Cache**: PostgreSQL 16 (JSONB, pgvector) + Redis 7
- **실시간 진행률 피드백**: 각 에이전트의 상태와 단계별 진행 상황을 WebSocket으로 실시간 시각화

---

## 🏗️ 시스템 아키텍처

```
[사용자]
    ↓
[Nuxt.js Frontend (포트 3000)]  ←(WebSocket / REST)→  [Go Backend API (포트 8080)]
                                                               ↓
                                                     [PostgreSQL + Redis]
                                                               ↓
                                                     [FastAPI AI Server (포트 8000)]
                                                               ↓
                                                     [CrewAI / LangGraph]
                                                          /    |    \
                                                   [Search] [Verify] [Writer]
```

---

## 📁 프로젝트 구조

```
ResearchForge/
├── .AI_Handover/            # AI 에이전트별 인수인계 및 실행 가이드
├── Docs/                    # 프로젝트 공식 문서 체계
│   ├── plan/                # 기획서, 인수인계서, 명세서
│   ├── Architecture/        # 아키텍처 및 데이터 흐름 문서
│   ├── branchs/commits/     # 커밋별 작업 기록 문서 (SSOT 준수)
│   ├── find_debug/          # 디버깅 및 문제 해결 기록
│   ├── guidance/            # 실행 및 개발 환경 가이드
│   ├── trainings/           # 연구 및 학습 자료
│   ├── repeater/            # 인프라/네트워크 장애 및 복구 이력
│   ├── vaildates/           # 기능 검증 및 테스트 체크리스트
│   └── web_management/      # 웹 서비스 및 배포 관리 기록
├── frontend/                # Nuxt.js 3 웹 애플리케이션
├── backend/                 # Go 백엔드 API & WebSocket 서버
├── ai-server/               # Python FastAPI 멀티 에이전트 오케스트레이터
├── scripts/                 # DB 초기화(init.sql) 및 유틸리티
├── docker-compose.yml       # 통합 개발/배포 환경 정의
├── AGENTS.md                # [SSOT] 모든 에이전트 공통 규칙
├── CONTRIBUTING.md          # 기여 가이드
├── LICENSE                  # MIT License
└── SECURITY.md              # 보안 정책
```

---

## 🚀 빠른 시작 (Docker Compose)

### 1. 환경 변수 설정
```bash
cp .env.example .env
# .env 파일에서 LLM API 키(OpenAI, Anthropic 등) 및 검색 API 키(Tavily 등)를 설정합니다.
```

### 2. 컨테이너 일괄 빌드 및 실행
```bash
docker-compose up -d --build
```

### 3. 접속 주소
- **Frontend Web UI**: `http://localhost:3000`
- **Go Backend API**: `http://localhost:8080`
- **FastAPI AI Server Docs**: `http://localhost:8000/docs`

---

## 📜 코딩 및 커밋 규칙

모든 개발 및 AI 에이전트 작업은 루트의 `AGENTS.md`를 단일 진실 공급원(SSOT)으로 삼습니다:
- **작업 브랜치**: `feature/*` 브랜치에서만 작업 (`main`/`develop`은 병합 전용)
- **커밋 메시지**: `type(scope): 한국어 요약`
- **커밋 기록 문서**: 매 커밋 시 `Docs/branchs/commits/YYMMDD_HHMM_[id]_[이름]_[브랜치].md` 작성 및 동시 커밋

---

## 📄 라이선스

이 프로젝트는 [MIT License](LICENSE)에 따라 배포됩니다.
