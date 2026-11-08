<script setup lang="ts">
import { ref } from 'vue'
import type { Report } from '~/types'
import SourceCard from '~/components/SourceCard.vue'

const props = defineProps<{
  report: Report
}>()

const activeTab = ref<'content' | 'sources'>('content')

const handleCopy = () => {
  if (navigator.clipboard) {
    navigator.clipboard.writeText(props.report.contentMarkdown || props.report.content_markdown)
    alert('보고서 마크다운이 클립보드에 복사되었습니다.')
  }
}
</script>

<template>
  <div class="bg-slate-900 border border-slate-800 rounded-2xl p-6 sm:p-8 shadow-2xl">
    <!-- 헤더 컨트롤 -->
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 pb-6 border-b border-slate-800">
      <div class="flex items-center gap-3">
        <button
          @click="activeTab = 'content'"
          :class="activeTab === 'content' ? 'bg-indigo-600 text-white shadow-md' : 'text-slate-400 hover:text-slate-200'"
          class="px-4 py-2 rounded-lg text-sm font-semibold transition-all cursor-pointer"
        >
          📄 보고서 본문
        </button>
        <button
          @click="activeTab = 'sources'"
          :class="activeTab === 'sources' ? 'bg-indigo-600 text-white shadow-md' : 'text-slate-400 hover:text-slate-200'"
          class="px-4 py-2 rounded-lg text-sm font-semibold transition-all cursor-pointer flex items-center gap-1.5"
        >
          <span>🔗 검증된 출처</span>
          <span class="bg-indigo-900/80 text-indigo-300 text-xs px-2 py-0.5 rounded-full">
            {{ report.sources?.length || 0 }}
          </span>
        </button>
      </div>

      <div class="flex items-center gap-2">
        <button
          @click="handleCopy"
          class="px-3.5 py-1.5 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-200 text-xs font-medium transition-all cursor-pointer"
        >
          📋 Markdown 복사
        </button>
      </div>
    </div>

    <!-- 탭 1: 보고서 본문 -->
    <div v-if="activeTab === 'content'" class="pt-6">
      <div class="prose prose-invert max-w-none prose-headings:text-slate-100 prose-p:text-slate-300 prose-a:text-indigo-400 prose-blockquote:border-l-indigo-500 whitespace-pre-wrap font-sans text-sm sm:text-base leading-relaxed">
        {{ report.contentMarkdown || report.content_markdown }}
      </div>
    </div>

    <!-- 탭 2: 검증된 출처 목록 -->
    <div v-else-if="activeTab === 'sources'" class="pt-6 space-y-3">
      <div v-if="!report.sources || report.sources.length === 0" class="text-center py-10 text-slate-500">
        등록된 출처 정보가 없습니다.
      </div>
      <SourceCard
        v-for="(source, idx) in report.sources"
        :key="idx"
        :source="source"
      />
    </div>
  </div>
</template>
