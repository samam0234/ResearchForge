import asyncio
from datetime import datetime
from typing import List
from app.models.schemas import RunCrewRequest, RunCrewResponse, ReportSchema, AgentLog, SourceItem
from app.agents.researcher import ResearcherAgent
from app.agents.fact_checker import FactCheckerAgent
from app.agents.analyst import AnalystAgent
from app.agents.writer import WriterAgent

class ResearchCrew:
    """CrewAI 기반 멀티 에이전트 리서치 크루 오케스트레이터"""

    def __init__(self, request: RunCrewRequest):
        self.request = request
        self.logs: List[AgentLog] = []

    async def kickoff(self) -> RunCrewResponse:
        project_id = self.request.project_id
        query = self.request.query
        options = self.request.options

        try:
            # Step 1: Researcher 실행
            start_t = datetime.utcnow()
            raw_data = await ResearcherAgent.run(query, min_sources=options.min_sources)
            self.logs.append(AgentLog(
                project_id=project_id,
                agent_name="Researcher",
                role=ResearcherAgent.role,
                status="completed",
                started_at=start_t,
                finished_at=datetime.utcnow()
            ))

            # Step 2: FactChecker 실행
            start_t = datetime.utcnow()
            verified_data = await FactCheckerAgent.run(raw_data)
            self.logs.append(AgentLog(
                project_id=project_id,
                agent_name="FactChecker",
                role=FactCheckerAgent.role,
                status="completed",
                started_at=start_t,
                finished_at=datetime.utcnow()
            ))

            # Step 3: Analyst 실행
            start_t = datetime.utcnow()
            analysis_data = await AnalystAgent.run(query, verified_data)
            self.logs.append(AgentLog(
                project_id=project_id,
                agent_name="Analyst",
                role=AnalystAgent.role,
                status="completed",
                started_at=start_t,
                finished_at=datetime.utcnow()
            ))

            # Step 4: Writer 실행
            start_t = datetime.utcnow()
            sources: List[SourceItem] = verified_data.get("verified_sources", [])
            report_md = await WriterAgent.run(query, analysis_data, sources)
            self.logs.append(AgentLog(
                project_id=project_id,
                agent_name="Writer",
                role=WriterAgent.role,
                status="completed",
                started_at=start_t,
                finished_at=datetime.utcnow()
            ))

            final_report = ReportSchema(
                project_id=project_id,
                content_markdown=report_md,
                sources=sources,
                metadata={
                    "total_sources": len(sources),
                    "depth": options.depth,
                    "language": options.language,
                    "pipeline": "CrewAI-Sequential"
                },
                created_at=datetime.utcnow()
            )

            return RunCrewResponse(
                project_id=project_id,
                status="completed",
                report=final_report,
                logs=self.logs
            )

        except Exception as e:
            return RunCrewResponse(
                project_id=project_id,
                status="failed",
                error=str(e),
                logs=self.logs
            )
