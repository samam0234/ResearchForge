from typing import Dict, Any, List
from app.tools.search import SearchTool

class ResearcherAgent:
    """웹 검색 및 방대한 정보 수집을 담당하는 에이전트"""
    role = "Researcher (연구 조사관)"
    goal = "사용자 쿼리에 부합하는 신뢰도 높은 최신 웹 문서와 원천 자료를 광범위하게 수집한다."
    backstory = "당신은 10년 이상의 학술 및 시장 조사 경험을 가진 시니어 리서처입니다. 노이즈를 필터링하고 핵심 팩트가 담긴 출처를 찾아내는 데 탁월합니다."

    @classmethod
    async def run(cls, query: str, min_sources: int = 5) -> Dict[str, Any]:
        results = await SearchTool.search(query, max_results=min_sources)
        return {
            "agent": "Researcher",
            "query": query,
            "raw_sources": results,
            "count": len(results)
        }
