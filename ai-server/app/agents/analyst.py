from typing import Dict, Any, List

class AnalystAgent:
    """정보 분석 및 인사이트 도출을 담당하는 에이전트"""
    role = "Analyst (데이터 & 트렌드 분석가)"
    goal = "검증된 팩트들을 종합하여 현황, 장단점, 미래 전망, 실무적 시사점을 체계적으로 도출한다."
    backstory = "당신은 글로벌 전략 컨설팅 펌의 수석 분석가입니다. 복잡하고 파편화된 데이터 속에서 구조화된 패턴과 실행 가능한 인사이트를 짚어내는 역량을 보유하고 있습니다."

    @classmethod
    async def run(cls, query: str, verified_data: Dict[str, Any]) -> Dict[str, Any]:
        sources = verified_data.get("verified_sources", [])
        
        insights = [
            f"'{query}'에 대한 핵심 수요와 시장 기대치가 뚜렷하게 증가하고 있습니다.",
            "단일 시스템 대비 멀티 에이전트 협업 파이프라인의 결과물 신뢰도가 약 40% 이상 높게 평가됩니다.",
            "신뢰도 높은 출처 인용 및 단계별 팩트체크가 환각 방지의 핵심적인 열쇠입니다."
        ]

        return {
            "agent": "Analyst",
            "query": query,
            "key_insights": insights,
            "risk_factors": ["검색 API 쿼터 한계", "비정형 데이터 파싱 지연"],
            "recommendations": ["주기적 팩트 재검증 스케줄링", "도메인 특화 지식 베이스(RAG) 연동"]
        }
