<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useResearch } from '~/composables/useResearch'
import { useWebSocket } from '~/composables/useWebSocket'
import AgentTimeline from '~/components/AgentTimeline.vue'
import ReportViewer from '~/components/ReportViewer.vue'
import type { ResearchDetail } from '~/types'

const route = useRoute()
const projectId = route.params.id as string

const { getResearchDetail, exportReport } = useResearch()
const { events, currentStep, progress, isFinished, connect } = useWebSocket(projectId)

const detail = ref<ResearchDetail | null>(null)
const loading = ref(true)

const fetchDetail = async () => {
  try {
    detail.value = await getResearchDetail(projectId)
  } catch (e) {
    console.error('Failed to load research detail:', e)
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  await fetchDetail()
  connect()
})

// WebSocket 완료 이벤트 수신 시 최신 보고서 데이터 갱신
watch(isFinished, async (finished) => {
  if (finished) {
    await fetchDetail()
  }
})
</script>

<template>
  <div class="space-y-8">
    <!-- 상단 헤더 정보 -->
    <div v-if="detail" class="bg-slate-900 border border-slate-800 rounded-2xl p-6 flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div class="space-y-1">
        <div class="flex items-center gap-2">
          <NuxtLink to="/history" class="text-xs text-slate-400 hover:text-slate-200">리서치 목록</NuxtLink>
          <span class="text-slate-600">/</span>
          <span class="text-xs text-indigo-400 font-mono">{{ projectId.substring(0, 8) }}</span>
        </div>
        <h1 class="text-xl font-bold text-slate-100">{{ detail.project.title }}</h1>
        <p class="text-xs text-slate-400">
          작성일시: {{ new Date(detail.project.created_at).toLocaleString('ko-KR') }} · 탐색 깊이: {{ detail.project.options.depth }}
        </p>
      </div>

      <div class="flex items-center gap-3">
        <button
          v-if="detail.report"
          @click="exportReport(projectId, 'markdown')"
          class="px-4 py-2 rounded-xl bg-slate-800 hover:bg-slate-700 text-slate-100 text-xs font-semibold transition-all cursor-pointer flex items-center gap-1.5"
        >
          <span>📥 Markdown 내보내기</span>
        </button>
      </div>
    </div>

    <!-- 에이전트 실시간 타임라인 -->
    <AgentTimeline
      :events="events"
      :current-step="currentStep"
      :progress="detail?.project.status === 'completed' ? 100 : progress"
      :status="detail?.project.status || 'running'"
    />

    <!-- 완성된 보고서 뷰어 -->
    <div v-if="detail?.report" class="space-y-4">
      <div class="flex items-center justify-between">
        <h2 class="text-lg font-bold text-slate-100 flex items-center gap-2">
          <span>📑</span>
          <span>생성된 리서치 보고서</span>
        </h2>
      </div>
      <ReportViewer :report="detail.report" />
    </div>

    <div v-else-if="!loading && detail?.project.status !== 'completed'" class="text-center py-12 bg-slate-900/30 border border-dashed border-slate-800 rounded-2xl">
      <div class="text-3xl animate-bounce mb-3">⚙️</div>
      <h3 class="text-sm font-semibold text-slate-300 mb-1">에이전트들이 보고서를 작성하고 있습니다.</h3>
      <p class="text-xs text-slate-500">진행률이 100%에 도달하면 이곳에 검증된 출처와 함께 보고서가 표시됩니다.</p>
    </div>
  </div>
</template>
