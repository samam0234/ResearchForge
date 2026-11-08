import type { ResearchProject, ResearchDetail, ResearchOptions } from '~/types'

export const useResearch = () => {
  const config = useRuntimeConfig()
  const apiBase = config.public.apiBase

  const createResearch = async (title: string, query: string, options: Partial<ResearchOptions>): Promise<ResearchProject> => {
    return await $fetch<ResearchProject>(`${apiBase}/research`, {
      method: 'POST',
      body: { title, query, options }
    })
  }

  const getResearchDetail = async (id: string): Promise<ResearchDetail> => {
    return await $fetch<ResearchDetail>(`${apiBase}/research/${id}`)
  }

  const listResearches = async (): Promise<ResearchProject[]> => {
    const res = await $fetch<{ items: ResearchProject[] }>(`${apiBase}/research`)
    return res.items || []
  }

  const exportReport = (id: string, format = 'markdown') => {
    window.open(`${apiBase}/research/${id}/export?format=${format}`, '_blank')
  }

  return {
    createResearch,
    getResearchDetail,
    listResearches,
    exportReport
  }
}
