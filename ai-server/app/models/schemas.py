from pydantic import BaseModel, Field
from typing import List, Optional, Dict, Any
from datetime import datetime

class ResearchOptions(BaseModel):
    depth: str = Field(default="normal", description="shallow, normal, deep")
    min_sources: int = Field(default=5, description="최소 수집 출처 개수")
    language: str = Field(default="ko", description="ko, en")
    report_format: str = Field(default="markdown", description="markdown, academic, summary")

class RunCrewRequest(BaseModel):
    project_id: str
    query: str
    options: ResearchOptions = Field(default_factory=ResearchOptions)

class SourceItem(BaseModel):
    title: str
    url: str
    snippet: str
    credibility: float = Field(default=0.9, ge=0.0, le=1.0)

class AgentLog(BaseModel):
    project_id: str
    agent_name: str
    role: str
    status: str
    started_at: datetime
    finished_at: Optional[datetime] = None

class ReportSchema(BaseModel):
    project_id: str
    content_markdown: str
    sources: List[SourceItem] = []
    metadata: Dict[str, Any] = {}
    created_at: datetime = Field(default_factory=datetime.utcnow)

class RunCrewResponse(BaseModel):
    project_id: str
    status: str # completed, failed
    report: Optional[ReportSchema] = None
    logs: List[AgentLog] = []
    error: Optional[str] = None
