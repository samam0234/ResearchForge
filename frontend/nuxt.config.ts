// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss'],
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080/api',
      wsBase: process.env.NUXT_PUBLIC_WS_BASE || 'ws://localhost:8080/ws',
    }
  },
  app: {
    head: {
      title: 'ResearchForge - AI 기반 멀티 에이전트 리서치 도우미',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: '다중 에이전트가 협업하여 정보 수집, 출처 검증, 분석, 보고서 작성을 자동 수행하는 리서치 플랫폼' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  }
})
