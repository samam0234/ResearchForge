import httpx
from bs4 import BeautifulSoup
from typing import Optional

class WebScraperTool:
    """웹 페이지 본문 텍스트 추출 도구"""

    @classmethod
    async def scrape(cls, url: str) -> Optional[str]:
        try:
            async with httpx.AsyncClient(timeout=10.0, follow_redirects=True) as client:
                headers = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) ResearchForge/1.0"}
                resp = await client.get(url, headers=headers)
                if resp.status_code == 200:
                    soup = BeautifulSoup(resp.text, "html.parser")
                    # 불필요한 태그 제거
                    for tag in soup(["script", "style", "nav", "footer", "header"]):
                        tag.decompose()
                    text = soup.get_text(separator=" ", strip=True)
                    return text[:3000] # 최대 3,000자 요약
        except Exception as e:
            print(f"[WebScraper] Scrape error for {url}: {e}")
        return None
