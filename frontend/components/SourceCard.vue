<script setup lang="ts">
import type { SourceItem } from '~/types'

defineProps<{
  source: SourceItem
}>()

const getCredibilityBadge = (score: number) => {
  if (score >= 0.9) return { text: '매우 높음', color: 'bg-emerald-500/20 text-emerald-400 border-emerald-500/30' }
  if (score >= 0.7) return { text: '신뢰 가능', color: 'bg-blue-500/20 text-blue-400 border-blue-500/30' }
  return { text: '참고용', color: 'bg-amber-500/20 text-amber-400 border-amber-500/30' }
}
</script>

<template>
  <div class="bg-slate-900/60 border border-slate-800 rounded-xl p-4 hover:border-slate-700 transition-all">
    <div class="flex items-start justify-between gap-3 mb-2">
      <h4 class="text-sm font-semibold text-slate-100 hover:text-indigo-400 line-clamp-1">
        <a :href="source.url" target="_blank" rel="noopener noreferrer" class="hover:underline">
          {{ source.title }}
        </a>
      </h4>
      <span
        :class="getCredibilityBadge(source.credibility).color"
        class="text-[11px] font-medium px-2 py-0.5 rounded-full border shrink-0"
      >
        신뢰도 {{ Math.round(source.credibility * 100) }}%
      </span>
    </div>
    <p class="text-xs text-slate-400 line-clamp-2 mb-2 leading-relaxed">
      {{ source.snippet }}
    </p>
    <a
      :href="source.url"
      target="_blank"
      rel="noopener noreferrer"
      class="text-[11px] text-indigo-400 hover:text-indigo-300 truncate block opacity-80 hover:opacity-100"
    >
      🔗 {{ source.url }}
    </a>
  </div>
</template>
