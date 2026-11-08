"""
LangGraph 기반 멀티 에이전트 상태 그래프 정의 (Phase 4 고도화용 스켈레톤)
- START -> research_node -> verify_node -> analyze_node -> write_node -> END
- 조건부 엣지(Conditional Edge)를 통한 재검증 루프 지원 설계
"""
from typing import TypedDict, List, Dict, Any, Optional
from app.models.schemas import SourceItem

class ResearchState(TypedDict):
    project_id: str
    query: str
    depth: str
    raw_sources: List[Dict[str, Any]]
    verified_sources: List[SourceItem]
    insights: List[str]
    report_markdown: Optional[str]
    retry_count: int
    is_valid: bool

class ResearchGraph:
    """LangGraph 호환 상태 그래프 정의"""

    @classmethod
    async def run(cls, project_id: str, query: str, depth: str = "normal") -> Dict[str, Any]:
        state: ResearchState = {
            "project_id": project_id,
            "query": query,
            "depth": depth,
            "raw_sources": [],
            "verified_sources": [],
            "insights": [],
            "report_markdown": None,
            "retry_count": 0,
            "is_valid": True
        }
        # 상태 노드 전이 시뮬레이션
        return state
