<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useResearch } from '~/composables/useResearch'
import type { ResearchProject } from '~/types'

const { listResearches } = useResearch()
const recentResearches = ref<ResearchProject[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    recentResearches.value = await listResearches()
  } catch (e) {
    console.warn('Could not fetch recent researches:', e)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="space-y-12">
    <!-- Hero Section -->
    <div class="text-center py-12 sm:py-16 space-y-6">
      <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-indigo-500/10 border border-indigo-500/20 text-indigo-400 text-xs font-semibold">
        <span>✨ 2026 차세대 멀티 에이전트 리서치 플랫폼</span>
      </div>
      <h1 class="text-3xl sm:text-5xl font-extrabold tracking-tight max-w-3xl mx-auto leading-tight">
        질문 하나로 시작하는 <br/>
        <span class="bg-gradient-to-r from-indigo-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">
          AI 심층 교차 리서치 보고서
        </span>
      </h1>
      <p class="text-slate-400 max-w-xl mx-auto text-base sm:text-lg">
        Researcher, FactChecker, Analyst, Writer 4개의 특화된 에이전트가 협업하여 팩트 기반의 구조화된 보고서를 자동으로 생성합니다.
      </p>

      <div class="flex items-center justify-center gap-4 pt-4">
        <NuxtLink
          to="/research/new"
          class="px-6 py-3.5 rounded-xl bg-gradient-to-r from-indigo-500 to-purple-600 hover:from-indigo-600 hover:to-purple-700 text-white font-semibold text-sm shadow-xl shadow-indigo-500/25 transition-all transform hover:-translate-y-0.5"
        >
          새 리서치 시작하기 🚀
        </NuxtLink>
        <NuxtLink
          to="/history"
          class="px-6 py-3.5 rounded-xl bg-slate-900 hover:bg-slate-800 border border-slate-800 text-slate-300 font-medium text-sm transition-all"
        >
          히스토리 둘러보기
        </NuxtLink>
      </div>
    </div>

    <!-- 4단계 에이전트 워크플로우 카드 -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-4">
      <div class="bg-slate-900/60 border border-slate-800 rounded-2xl p-5 space-y-2">
        <div class="text-2xl">🔍</div>
        <h3 class="font-bold text-slate-200 text-sm">1. 웹 검색 (Researcher)</h3>
        <p class="text-xs text-slate-400 leading-relaxed">
          Tavily/Serper 도구로 관련 문헌 및 최신 기사를 광범위하게 수집합니다.
        </p>
      </div>

      <div class="bg-slate-900/60 border border-slate-800 rounded-2xl p-5 space-y-2">
        <div class="text-2xl">🛡️</div>
        <h3 class="font-bold text-slate-200 text-sm">2. 팩트체크 (FactChecker)</h3>
        <p class="text-xs text-slate-400 leading-relaxed">
          수집된 원천 출처를 교차 검증하고 신뢰도 점수를 부여하여 환각을 억제합니다.
        </p>
      </div>

      <div class="bg-slate-900/60 border border-slate-800 rounded-2xl p-5 space-y-2">
        <div class="text-2xl">📊</div>
        <h3 class="font-bold text-slate-200 text-sm">3. 심층 분석 (Analyst)</h3>
        <p class="text-xs text-slate-400 leading-relaxed">
          검증된 데이터를 종합하여 패턴, 장단점, 핵심 인사이트를 구조화합니다.
        </p>
      </div>

      <div class="bg-slate-900/60 border border-slate-800 rounded-2xl p-5 space-y-2">
        <div class="text-2xl">✍️</div>
        <h3 class="font-bold text-slate-200 text-sm">4. 보고서 작성 (Writer)</h3>
        <p class="text-xs text-slate-400 leading-relaxed">
          인용 출처가 명시된 마크다운 전문 보고서를 정교하게 완성합니다.
        </p>
      </div>
    </div>

    <!-- 최근 리서치 목록 미리보기 -->
    <div class="space-y-4 pt-4">
      <div class="flex items-center justify-between">
        <h2 class="text-lg font-bold text-slate-200">최근 리서치 프로젝트</h2>
        <NuxtLink to="/history" class="text-xs text-indigo-400 hover:text-indigo-300 font-medium">전체보기 →</NuxtLink>
      </div>

      <div v-if="loading" class="text-center py-8 text-sm text-slate-500">
        데이터를 불러오는 중...
      </div>

      <div v-else-if="recentResearches.length === 0" class="text-center py-10 bg-slate-900/30 border border-dashed border-slate-800 rounded-2xl">
        <p class="text-sm text-slate-500 mb-3">아직 진행된 리서치가 없습니다.</p>
        <NuxtLink
          to="/research/new"
          class="text-xs text-indigo-400 hover:underline font-semibold"
        >
          첫 번째 리서치를 시작해보세요!
        </NuxtLink>
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <NuxtLink
          v-for="item in recentResearches.slice(0, 4)"
          :key="item.id"
          :to="`/research/${item.id}`"
          class="block bg-slate-900 border border-slate-800 hover:border-indigo-500/50 rounded-xl p-5 transition-all group"
        >
          <div class="flex items-start justify-between gap-2 mb-2">
            <h3 class="font-semibold text-sm text-slate-200 group-hover:text-indigo-300 line-clamp-1">
              {{ item.title }}
            </h3>
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
          <p class="text-xs text-slate-400 line-clamp-2 mb-3">
            {{ item.query }}
          </p>
          <span class="text-[11px] text-slate-500">
            {{ new Date(item.created_at).toLocaleString('ko-KR') }}
          </span>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>
