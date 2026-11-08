<script setup lang="ts">
import { computed } from 'vue'
import type { WebSocketEvent } from '~/types'

const props = defineProps<{
  events: WebSocketEvent[]
  currentStep: string
  progress: number
  status: string
}>()

const agentSteps = [
  { name: 'Researcher', label: '정보 검색 및 수집', icon: '🔍' },
  { name: 'FactChecker', label: '출처 팩트체크', icon: '🛡️' },
  { name: 'Analyst', label: '인사이트 분석', icon: '📊' },
  { name: 'Writer', label: '보고서 작성', icon: '✍️' },
]

const getStepStatus = (stepName: string) => {
  if (props.status === 'completed') return 'completed'
  if (props.currentStep === stepName) return 'active'

  const stepOrder = ['Researcher', 'FactChecker', 'Analyst', 'Writer']
  const currentIndex = stepOrder.indexOf(props.currentStep)
  const thisIndex = stepOrder.indexOf(stepName)

  if (currentIndex > thisIndex) return 'completed'
  return 'pending'
}
</script>

<template>
  <div class="bg-slate-900 border border-slate-800 rounded-2xl p-6 shadow-xl">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-base font-bold text-slate-100 flex items-center gap-2">
        <span>🤖</span>
        <span>멀티 에이전트 협업 파이프라인</span>
      </h3>
      <span class="text-xs font-semibold text-indigo-400 bg-indigo-500/10 px-2.5 py-1 rounded-full border border-indigo-500/20">
        진행률 {{ progress }}%
      </span>
    </div>

    <!-- 진행률 바 -->
    <div class="w-full bg-slate-950 rounded-full h-2.5 mb-6 overflow-hidden">
      <div
        class="bg-gradient-to-r from-indigo-500 via-purple-500 to-emerald-400 h-2.5 rounded-full transition-all duration-500 ease-out"
        :style="{ width: `${progress}%` }"
      ></div>
    </div>

    <!-- 에이전트 단계 카드 -->
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-3 mb-6">
      <div
        v-for="step in agentSteps"
        :key="step.name"
        :class="{
          'border-indigo-500 bg-indigo-950/30 text-white ring-1 ring-indigo-500/50': getStepStatus(step.name) === 'active',
          'border-emerald-500/50 bg-emerald-950/20 text-emerald-300': getStepStatus(step.name) === 'completed',
          'border-slate-800 bg-slate-950/50 text-slate-500': getStepStatus(step.name) === 'pending'
        }"
        class="border rounded-xl p-3 flex flex-col justify-between transition-all"
      >
        <div class="flex items-center justify-between mb-1">
          <span class="text-lg">{{ step.icon }}</span>
          <span
            v-if="getStepStatus(step.name) === 'active'"
            class="animate-pulse inline-block w-2 h-2 rounded-full bg-indigo-400"
          ></span>
          <span
            v-else-if="getStepStatus(step.name) === 'completed'"
            class="text-xs text-emerald-400 font-bold"
          >✓</span>
        </div>
        <div>
          <div class="text-xs font-semibold">{{ step.name }}</div>
          <div class="text-[11px] opacity-80">{{ step.label }}</div>
        </div>
      </div>
    </div>

    <!-- 실시간 이벤트 스트림 -->
    <div class="bg-slate-950 rounded-xl p-3 max-h-48 overflow-y-auto font-mono text-xs text-slate-400 space-y-1.5 border border-slate-800/80">
      <div v-if="events.length === 0" class="text-center py-4 text-slate-600">
        에이전트 연결 대기 중...
      </div>
      <div
        v-for="(evt, idx) in events"
        :key="idx"
        class="flex items-start gap-2 text-[11px]"
      >
        <span class="text-slate-500 shrink-0">[{{ evt.agent_name || 'System' }}]</span>
        <span class="text-slate-300">{{ evt.message }}</span>
      </div>
    </div>
  </div>
</template>
