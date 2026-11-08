<script setup lang="ts">
import { ref } from 'vue'
import type { ResearchOptions } from '~/types'

const emit = defineEmits<{
  (e: 'submit', payload: { title: string; query: string; options: ResearchOptions }): void
}>()

const query = ref('')
const title = ref('')
const depth = ref<'shallow' | 'normal' | 'deep'>('normal')
const minSources = ref(5)
const language = ref('ko')
const reportFormat = ref('markdown')
const isSubmitting = ref(false)

const handleSubmit = () => {
  if (!query.value.trim()) return

  isSubmitting.value = true
  emit('submit', {
    title: title.value.trim() || query.value.trim(),
    query: query.value.trim(),
    options: {
      depth: depth.value,
      min_sources: minSources.value,
      language: language.value,
      report_format: reportFormat.value,
    }
  })
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="bg-slate-900 border border-slate-800 rounded-2xl p-6 sm:p-8 shadow-2xl backdrop-blur-sm">
    <div class="mb-6">
      <label for="query" class="block text-sm font-semibold text-slate-200 mb-2">
        연구 질문 또는 리서치 주제 <span class="text-indigo-400">*</span>
      </label>
      <textarea
        id="query"
        v-model="query"
        rows="4"
        required
        placeholder="예: 2026년 차세대 AI 멀티 에이전트 프레임워크 비교 및 시장 도입 현황을 조사해줘."
        class="w-full bg-slate-950/80 border border-slate-700 rounded-xl p-4 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all text-base"
      ></textarea>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
      <div>
        <label class="block text-xs font-medium text-slate-400 mb-1">탐색 깊이 (Research Depth)</label>
        <select
          v-model="depth"
          class="w-full bg-slate-950 border border-slate-700 rounded-lg p-2.5 text-sm text-slate-200 focus:ring-2 focus:ring-indigo-500"
        >
          <option value="shallow">신속 탐색 (빠른 핵심 요약)</option>
          <option value="normal">표준 탐색 (균형 잡힌 분석)</option>
          <option value="deep">심층 탐색 (포괄적 검증 & 인용)</option>
        </select>
      </div>

      <div>
        <label class="block text-xs font-medium text-slate-400 mb-1">최소 검증 출처 수</label>
        <input
          type="number"
          v-model.number="minSources"
          min="3"
          max="20"
          class="w-full bg-slate-950 border border-slate-700 rounded-lg p-2.5 text-sm text-slate-200 focus:ring-2 focus:ring-indigo-500"
        />
      </div>

      <div>
        <label class="block text-xs font-medium text-slate-400 mb-1">보고서 언어</label>
        <select
          v-model="language"
          class="w-full bg-slate-950 border border-slate-700 rounded-lg p-2.5 text-sm text-slate-200 focus:ring-2 focus:ring-indigo-500"
        >
          <option value="ko">한국어 (Korean)</option>
          <option value="en">영어 (English)</option>
        </select>
      </div>
    </div>

    <div class="flex items-center justify-between pt-4 border-t border-slate-800">
      <span class="text-xs text-slate-400">
        ⚡ 4개 에이전트(검색, 팩트체크, 분석, 작성)가 순차 협업합니다.
      </span>
      <button
        type="submit"
        :disabled="isSubmitting || !query.trim()"
        class="inline-flex items-center gap-2 px-6 py-3 rounded-xl bg-gradient-to-r from-indigo-500 to-purple-600 hover:from-indigo-600 hover:to-purple-700 text-white font-semibold text-sm transition-all shadow-lg shadow-indigo-500/25 disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
      >
        <span v-if="isSubmitting" class="animate-spin">🔄</span>
        <span>리서치 시작하기</span>
      </button>
    </div>
  </form>
</template>
