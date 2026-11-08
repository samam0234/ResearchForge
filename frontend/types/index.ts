export interface ResearchOptions {
  depth: 'shallow' | 'normal' | 'deep'
  min_sources: number
  language: string
  report_format: string
}

export interface ResearchProject {
  id: string
  user_id: string
  title: string
  query: string
  options: ResearchOptions
  status: 'pending' | 'running' | 'completed' | 'failed'
  created_at: string
  completed_at?: string
}

export interface AgentExecution {
  id: string
  project_id: string
  agent_name: string
  role: string
  status: 'running' | 'completed' | 'failed'
  started_at: string
  finished_at?: string
}

export interface SourceItem {
  title: string
  url: string
  snippet: string
  credibility: number
}

export interface Report {
  id: string
  project_id: string
  content_markdown: string
  sources: SourceItem[]
  metadata: Record<string, any>
  created_at: string
}

export interface ResearchDetail {
  project: ResearchProject
  executions: AgentExecution[]
  report?: Report
}

export interface WebSocketEvent {
  type: 'CONNECTED' | 'PROJECT_STATUS' | 'STEP_START' | 'STEP_COMPLETE' | 'PROJECT_DONE' | 'ERROR'
  project_id: string
  agent_name?: string
  progress?: number
  message: string
  data?: any
  timestamp?: string
}
