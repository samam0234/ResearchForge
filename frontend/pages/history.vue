<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useResearch } from '~/composables/useResearch'
import type { ResearchProject } from '~/types'

const { listResearches } = useResearch()
const researches = ref<ResearchProject[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    researches.value = await listResearches()
  } catch (e) {
    console.error('Failed to load history:', e)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-100">리서치 히스토리</h1>
        <p class="text-xs text-slate-400">과거에 진행했던 모든 리서치 보고서 이력을 확인합니다.</p>
      </div>

      <NuxtLink
        to="/research/new"
        class="px-4 py-2 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-xs shadow-md shadow-indigo-500/20 transition-all"
      >
        + 새 리서치
      </NuxtLink>
    </div>

    <div v-if="loading" class="text-center py-12 text-sm text-slate-500">
      히스토리 목록을 불러오는 중...
    </div>

    <div v-else-if="researches.length === 0" class="text-center py-16 bg-slate-900/30 border border-dashed border-slate-800 rounded-2xl">
      <div class="text-4xl mb-3">📁</div>
      <p class="text-sm font-semibold text-slate-400 mb-1">저장된 리서치 기록이 없습니다.</p>
      <p class="text-xs text-slate-500 mb-4">궁금한 주제를 입력하여 새로운 리서치를 시작해 보세요.</p>
      <NuxtLink
        to="/research/new"
        class="px-4 py-2 rounded-lg bg-indigo-600 text-white text-xs font-semibold hover:bg-indigo-500"
      >
        첫 리서치 작성하기
      </NuxtLink>
    </div>

    <div v-else class="space-y-3">
      <NuxtLink
        v-for="item in researches"
        :key="item.id"
        :to="`/research/${item.id}`"
        class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 bg-slate-900 border border-slate-800 hover:border-indigo-500/50 rounded-xl p-5 transition-all group"
      >
        <div class="space-y-1">
          <div class="flex items-center gap-2">
            <span class="text-xs font-mono text-slate-500">{{ item.id.substring(0, 8) }}</span>
            <span
              :class="{
                'bg-emerald-500/20 text-emerald-400 border-emerald-500/30': item.status === 'completed',
                'bg-indigo-500/20 text-indigo-400 border-indigo-500/30 animate-pulse': item.status === 'running',
                'bg-slate-700/30 text-slate-400 border-slate-600': item.status === 'pending',
                'bg-rose-500/20 text-rose-400 border-rose-500/30': item.status === 'failed',
              }"
              class="text-[10px] font-semibold px-2 py-0.5 rounded-full border shrink-0"
            >
              {{ item.status }}
            </span>
          </div>
          <h3 class="text-base font-semibold text-slate-200 group-hover:text-indigo-300">
            {{ item.title }}
          </h3>
          <p class="text-xs text-slate-400 line-clamp-1">
            {{ item.query }}
          </p>
        </div>

        <div class="flex items-center gap-4 text-xs text-slate-500 shrink-0">
          <span>{{ new Date(item.created_at).toLocaleDateString('ko-KR') }}</span>
          <span class="text-indigo-400 group-hover:translate-x-1 transition-transform">→</span>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>
