# ResearchForge 기획서 v1.0

**AI 기반 멀티 에이전트 리서치 도우미**

- 문서 버전: v1.0
- 작성일: 2026-11-09
- 프로젝트 상태: 기획 단계
- 문서 유형: 프로젝트 명세서 / 개발 계획서

---

## 1. 프로젝트 개요

### 1.1 프로젝트명
**ResearchForge**

### 1.2 한 줄 소개
사용자가 연구 주제나 질문만 입력하면, 여러 AI 에이전트가 협업하여 정보 수집 · 검증 · 분석 · 보고서 작성을 자동으로 수행하는 멀티 에이전트 리서치 플랫폼.

### 1.3 배경 및 필요성
- 정보 과잉 시대에 단순 검색만으로는 신뢰할 수 있는 리서치가 어려움
- 단일 LLM 응답은 편향·환각·출처 부족 문제가 잦음
- 여러 역할(검색, 검증, 분석, 작성)을 가진 에이전트가 협업하면 품질과 신뢰도가 크게 향상됨
- 학습·업무·콘텐츠 제작 등에서 반복되는 리서치 작업을 자동화할 필요가 있음

### 1.4 핵심 가치
- **신뢰성**: 출처 기반 검증과 다중 에이전트 교차 검증
- **효율성**: 주제 입력 → 구조화된 보고서까지 자동화
- **확장성**: 에이전트 역할과 도구를 쉽게 추가 가능
- **투명성**: 각 에이전트의 작업 과정과 근거를 사용자에게 노출

---

## 2. 프로젝트 목표

### 2.1 주요 목표
1. 사용자가 자연어로 리서치 요청을 하면 멀티 에이전트가 협업하여 구조화된 보고서를 생성한다.
2. 각 에이전트의 역할(검색, 팩트체크, 분석, 작성)을 명확히 분리하고 워크플로우를 시각화한다.
3. 프론트엔드(Nuxt.js), 백엔드(Go), AI 서버(FastAPI)가 명확히 분리된 클린 아키텍처를 구현한다.
4. CrewAI 또는 LangGraph를 활용한 에이전트 오케스트레이션을 학습·적용한다.

### 2.2 성공 지표 (MVP 기준)
- 리서치 요청 후 평균 응답 시간 3분 이내
- 보고서에 최소 5개 이상의 출처 포함
- 사용자 만족도(자체 테스트) 4점 이상 / 5점
- 에이전트 워크플로우 실패율 10% 이하

---

## 3. 기술 스택

| 구분 | 기술 | 선정 이유 |
|------|------|-----------|
| **Frontend** | Nuxt.js 3 (Vue 3) | SSR/SSG 지원, 파일 기반 라우팅, 개발 경험 좋음 |
| **Backend** | Go (Fiber 또는 Chi) | 고성능, 동시성 처리 우수, 타입 안정성 |
| **AI Server** | FastAPI (Python) | AI 생태계와 궁합 좋음, 비동기 지원 |
| **AI Layer** | CrewAI 또는 LangGraph | 멀티 에이전트 오케스트레이션 특화 |
| **Database** | PostgreSQL | 관계형 + JSONB 지원, pgvector 확장 가능 |
| **Vector DB** | pgvector 또는 Qdrant | RAG 및 유사 문서 검색용 |
| **Cache / Queue** | Redis | 세션, 작업 큐, 실시간 상태 관리 |
| **Search Tool** | Tavily / Serper / DuckDuckGo | 웹 검색 도구 |
| **Authentication** | JWT + OAuth2 | 사용자 인증 |
| **Deployment** | Docker Compose (MVP) → K8s (향후) | 개발·배포 편의성 |

### 3.1 AI 레이어 선택 가이드
- **CrewAI**: 역할 기반 에이전트 정의가 직관적이고 빠르게 프로토타이핑 가능. MVP에 추천.
- **LangGraph**: 상태 머신 기반으로 복잡한 분기·루프·휴먼인더루프에 강함. 고도화 시 전환 고려.

---

## 4. 시스템 아키텍처

### 4.1 전체 구조

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

### 4.2 컴포넌트 역할

| 컴포넌트 | 주요 역할 |
|----------|-----------|
| **Nuxt.js** | UI/UX, 리서치 요청, 진행 상태 표시, 결과 리포트 뷰어 |
| **Go Backend** | 인증, 사용자 관리, 리서치 작업 큐잉, 상태 관리, AI 서버 오케스트레이션 |
| **FastAPI** | 에이전트 실행, 도구 호출, 결과 반환 |
| **PostgreSQL** | 사용자, 프로젝트, 리서치 히스토리, 보고서 저장 |
| **Redis** | 작업 상태, 실시간 진행률, 세션 |

### 4.3 통신 방식
- Frontend ↔ Backend: REST + WebSocket (진행 상태 실시간 업데이트)
- Backend ↔ AI Server: 내부 REST (gRPC 고려 가능)
- AI Server 내부: CrewAI/LangGraph 워크플로우

---

## 5. 주요 기능

### 5.1 핵심 기능 (MVP)

#### 5.1.1 리서치 요청
- 자연어 주제/질문 입력
- 옵션 설정: 깊이(얕음/보통/깊음), 언어, 출처 제한, 보고서 형식

#### 5.1.2 멀티 에이전트 협업
| 에이전트 | 역할 | 주요 도구 |
|----------|------|-----------|
| **Researcher** | 웹 검색 및 정보 수집 | Search Tool, Web Scraper |
| **FactChecker** | 출처 검증 및 교차 확인 | Search, Citation Extractor |
| **Analyst** | 정보 분석 및 인사이트 도출 | LLM Reasoning |
| **Writer** | 구조화된 보고서 작성 | Template Engine |
| **Critic** (선택) | 최종 검토 및 품질 개선 | Self-reflection |

#### 5.1.3 진행 상태 실시간 표시
- 각 에이전트별 현재 작업 표시
- 진행률 바 / 타임라인
- 중간 결과 미리보기

#### 5.1.4 보고서 뷰어
- Markdown 기반 보고서
- 출처 하이라이트 및 링크
- 섹션별 접기/펼치기
- PDF/Markdown 내보내기

#### 5.1.5 히스토리 관리
- 과거 리서치 목록
- 재실행 / 수정 요청
- 즐겨찾기

### 5.2 부가 기능 (v1.1 이후)
- 프로젝트 단위 리서치 그룹화
- 팀 공유 및 협업 코멘트
- 사용자 맞춤 에이전트 프롬프트 튜닝
- 로컬 문서 업로드 후 RAG 결합
- 음성 입력 지원

---

## 6. 사용자 시나리오

### 시나리오 1: 기본 리서치
1. 사용자가 "2026년 AI 에이전트 시장 동향" 입력
2. 옵션에서 "깊이: 보통", "출처 5개 이상" 선택
3. 에이전트들이 순차/병렬로 작업 수행
4. 3~5분 후 구조화된 보고서 생성
5. 사용자가 특정 섹션을 클릭해 상세 근거 확인

### 시나리오 2: 반복 리서치
1. 이전 리서치 결과를 기반으로 "최신 업데이트만 반영" 요청
2. 기존 컨텍스트를 유지한 채 증분 리서치 수행

---

## 7. 데이터 모델 및 DB 설계

### 7.1 주요 엔티티

```sql
-- 사용자
users (
  id UUID PRIMARY KEY,
  email VARCHAR UNIQUE,
  name VARCHAR,
  created_at TIMESTAMP
)

-- 리서치 프로젝트
research_projects (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users,
  title VARCHAR,
  query TEXT,
  options JSONB,
  status VARCHAR, -- pending, running, completed, failed
  created_at TIMESTAMP,
  completed_at TIMESTAMP
)

-- 에이전트 실행 로그
agent_executions (
  id UUID PRIMARY KEY,
  project_id UUID REFERENCES research_projects,
  agent_name VARCHAR,
  role VARCHAR,
  input JSONB,
  output JSONB,
  status VARCHAR,
  started_at TIMESTAMP,
  finished_at TIMESTAMP
)

-- 최종 보고서
reports (
  id UUID PRIMARY KEY,
  project_id UUID REFERENCES research_projects,
  content_markdown TEXT,
  sources JSONB,
  metadata JSONB,
  created_at TIMESTAMP
)
```

### 7.2 인덱스 전략
- `research_projects(user_id, created_at DESC)`
- `agent_executions(project_id)`
- pgvector 확장 시 임베딩 컬럼 추가 고려

---

## 8. API 설계 (주요 엔드포인트)

### 8.1 Go Backend API

| Method | Endpoint | 설명 |
|--------|----------|------|
| POST | `/api/auth/login` | 로그인 |
| POST | `/api/research` | 리서치 요청 생성 |
| GET | `/api/research/:id` | 리서치 상태 및 결과 조회 |
| GET | `/api/research` | 사용자 리서치 목록 |
| WS | `/ws/research/:id` | 실시간 진행 상태 |
| POST | `/api/research/:id/export` | 보고서 내보내기 |

### 8.2 FastAPI AI Server

| Method | Endpoint | 설명 |
|--------|----------|------|
| POST | `/ai/crew/run` | CrewAI 워크플로우 실행 |
| POST | `/ai/graph/run` | LangGraph 워크플로우 실행 |
| GET | `/ai/status/:task_id` | 작업 상태 조회 |

---

## 9. 에이전트 워크플로우

### 9.1 CrewAI 기반 (MVP 추천)

```
Researcher → FactChecker → Analyst → Writer → (Critic)
```

- **순차 실행**을 기본으로 하되, Researcher는 병렬 검색 가능
- 각 에이전트는 명확한 `role`, `goal`, `backstory` 정의
- 도구: `TavilySearch`, `WebScraper`, `CitationValidator`

### 9.2 LangGraph 기반 (고도화)

```
START → research_node → verify_node → analyze_node → write_node → END
                ↑___________________________|
                     (필요 시 재검색 루프)
```

- 조건부 엣지로 품질이 낮을 경우 재시도
- 상태(State)에 중간 결과를 누적

### 9.3 프롬프트 설계 원칙
- 각 에이전트는 자신의 역할에만 집중
- 출력 형식은 구조화된 JSON 또는 Markdown으로 강제
- 출처는 반드시 URL + 인용문 포함

---

## 10. 프론트엔드 구조 (Nuxt.js)

```
pages/
  index.vue                 # 랜딩 / 대시보드
  research/
    new.vue                 # 새 리서치
    [id].vue                # 진행/결과 페이지
  history.vue               # 히스토리
components/
  ResearchForm.vue
  AgentTimeline.vue
  ReportViewer.vue
  SourceCard.vue
composables/
  useResearch.ts
  useWebSocket.ts
```

### 주요 UI 요소
- 리서치 입력 폼 (주제 + 옵션)
- 에이전트 타임라인 (실시간)
- Markdown 보고서 렌더러 (출처 하이라이트)
- 진행률 및 예상 소요 시간

---

## 11. 백엔드 구조 (Go)

```
cmd/server/main.go
internal/
  handler/          # HTTP 핸들러
  service/          # 비즈니스 로직
  repository/       # DB 접근
  client/           # AI 서버 클라이언트
  model/            # 도메인 모델
  middleware/       # 인증, 로깅
pkg/
  queue/            # Redis 기반 작업 큐
```

### 주요 책임
- 사용자 인증 및 권한
- 리서치 작업 생성 및 상태 관리
- AI 서버 호출 및 결과 저장
- WebSocket을 통한 실시간 푸시

---

## 12. AI 서버 구조 (FastAPI)

```
app/
  main.py
  crews/
    research_crew.py      # CrewAI 정의
  graphs/
    research_graph.py     # LangGraph 정의
  agents/
    researcher.py
    fact_checker.py
    analyst.py
    writer.py
  tools/
    search.py
    scraper.py
  models/
    schemas.py
```

---

## 13. 개발 단계 및 마일스톤

| 단계 | 기간 | 목표 | 산출물 |
|------|------|------|--------|
| **Phase 0** | 1주 | 환경 구축 및 프로토타입 | Docker Compose, 기본 에이전트 동작 |
| **Phase 1** | 2주 | MVP 핵심 기능 | 리서치 요청 → 보고서 생성 E2E |
| **Phase 2** | 2주 | 실시간 UI + 히스토리 | WebSocket, 보고서 뷰어, 저장 |
| **Phase 3** | 1~2주 | 품질 개선 및 배포 | 에러 핸들링, 테스트, 배포 |
| **Phase 4** | 이후 | 고도화 | LangGraph 전환, RAG, 팀 기능 |

### MVP 완료 기준
- [ ] 주제 입력 후 보고서 생성 성공
- [ ] 최소 3개 에이전트 협업 확인
- [ ] 출처 5개 이상 포함
- [ ] 프론트에서 결과 확인 가능

---

## 14. 비기능 요구사항

| 항목 | 요구사항 |
|------|----------|
| **성능** | 리서치 평균 3분 이내 (깊이 보통 기준) |
| **가용성** | MVP 기준 99% (단일 인스턴스) |
| **보안** | JWT 인증, API 키 암호화 저장, 입력 검증 |
| **확장성** | AI 서버 수평 확장 가능 구조 |
| **관찰성** | 구조화된 로깅, 에이전트별 트레이스 |
| **비용** | 로컬 LLM(Ollama) 옵션 지원으로 API 비용 절감 |

---

## 15. 리스크 및 대응 방안

| 리스크 | 영향 | 대응 방안 |
|--------|------|-----------|
| LLM 환각 | 높음 | FactChecker 에이전트 + 출처 강제 |
| 검색 품질 저하 | 중간 | 다중 검색 도구 + 재시도 로직 |
| 응답 시간 초과 | 중간 | 타임아웃 + 부분 결과 반환 |
| API 비용 증가 | 중간 | 로컬 모델 옵션, 캐싱 |
| 에이전트 루프 | 낮음 | 최대 반복 횟수 제한 |

---

## 16. 향후 확장 계획

1. **로컬 문서 RAG 결합**: 사용자가 업로드한 PDF/문서를 컨텍스트로 활용
2. **멀티모달**: 이미지·차트 포함 보고서
3. **팀 협업**: 공유 프로젝트, 코멘트, 버전 관리
4. **에이전트 마켓플레이스**: 사용자 정의 에이전트 공유
5. **스케줄링**: 주기적 리서치 자동 실행
6. **LangGraph 완전 전환**: 복잡한 분기와 휴먼 인더 루프

---

## 17. 참고 및 학습 포인트

- Nuxt 3 Composition API + Server Routes
- Go Fiber/Chi + 구조화된 에러 처리
- FastAPI + Pydantic v2
- CrewAI 에이전트/태스크/크루 설계
- LangGraph StateGraph와 조건부 엣지
- WebSocket 실시간 통신
- PostgreSQL JSONB + pgvector

---

**문서 끝**

이 기획서는 개발 착수 전 기준으로 작성되었으며, 실제 구현 과정에서 세부 사항은 조정될 수 있습니다.
