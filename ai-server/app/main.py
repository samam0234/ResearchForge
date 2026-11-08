from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from app.models.schemas import RunCrewRequest, RunCrewResponse
from app.crews.research_crew import ResearchCrew
from app.graphs.research_graph import ResearchGraph
from app.config import settings

app = FastAPI(
    title=settings.app_name,
    version="1.0.0",
    description="ResearchForge AI Multi-Agent Orchestration Server"
)

# CORS 설정
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/health")
async def health_check():
    return {
        "status": "ok",
        "service": "researchforge-ai-server",
        "environment": settings.app_env
    }

@app.post("/ai/crew/run", response_model=RunCrewResponse)
async def run_crew(req: RunCrewRequest):
    """CrewAI 기반 멀티 에이전트 리서치 파이프라인 실행"""
    crew = ResearchCrew(request=req)
    result = await crew.kickoff()
    if result.status == "failed":
        raise HTTPException(status_code=500, detail=result.error or "Pipeline failed")
    return result

@app.post("/ai/graph/run")
async def run_graph(req: RunCrewRequest):
    """LangGraph 기반 멀티 에이전트 상태 그래프 파이프라인 실행 (고도화용)"""
    state = await ResearchGraph.run(
        project_id=req.project_id,
        query=req.query,
        depth=req.options.depth
    )
    return {"status": "completed", "state": state}

@app.get("/ai/status/{task_id}")
async def get_task_status(task_id: str):
    """에이전트 작업 상태 조회"""
    return {
        "task_id": task_id,
        "status": "completed"
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("app.main:app", host="0.0.0.0", port=settings.port, reload=True)
