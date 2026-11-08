<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useResearch } from '~/composables/useResearch'
import ResearchForm from '~/components/ResearchForm.vue'
import type { ResearchOptions } from '~/types'

const router = useRouter()
const { createResearch } = useResearch()

const handleStartResearch = async (payload: { title: string; query: string; options: ResearchOptions }) => {
  try {
    const project = await createResearch(payload.title, payload.query, payload.options)
    router.push(`/research/${project.id}`)
  } catch (err: any) {
    alert(`리서치 생성 중 오류가 발생했습니다: ${err.message || err}`)
  }
}
</script>

<template>
  <div class="max-w-3xl mx-auto space-y-6">
    <div class="space-y-1">
      <NuxtLink to="/" class="text-xs text-indigo-400 hover:underline">← 홈으로 돌아가기</NuxtLink>
      <h1 class="text-2xl font-bold text-slate-100">새로운 멀티 에이전트 리서치 생성</h1>
      <p class="text-sm text-slate-400">
        조사하고 싶은 주제나 기술 질문을 자연어로 상세히 기술해 주세요.
      </p>
    </div>

    <ResearchForm @submit="handleStartResearch" />
  </div>
</template>
