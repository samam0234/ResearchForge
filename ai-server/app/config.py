from pydantic_settings import BaseSettings
from typing import Optional

class Settings(BaseSettings):
    app_name: str = "ResearchForge AI Server"
    app_env: str = "development"
    port: int = 8000
    
    # LLM API Keys
    openai_api_key: Optional[str] = None
    anthropic_api_key: Optional[str] = None
    
    # Search API Keys
    tavily_api_key: Optional[str] = None
    serper_api_key: Optional[str] = None
    
    # Redis
    redis_host: str = "redis"
    redis_port: int = 6379

    class Config:
        env_file = ".env"
        extra = "allow"

settings = Settings()
