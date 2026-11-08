<script setup lang="ts">
// Global layout wrapper
</script>

<template>
  <div class="min-h-screen bg-slate-950 text-slate-100 flex flex-col font-sans">
    <!-- Header Navbar -->
    <header class="border-b border-slate-800/80 bg-slate-950/80 backdrop-blur-md sticky top-0 z-50">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 h-16 flex items-center justify-between">
        <NuxtLink to="/" class="flex items-center gap-2.5 font-bold text-lg tracking-tight group">
          <div class="w-8 h-8 rounded-lg bg-gradient-to-tr from-indigo-500 to-purple-600 flex items-center justify-center text-white shadow-md shadow-indigo-500/20 group-hover:scale-105 transition-transform">
            🔬
          </div>
          <span class="bg-gradient-to-r from-slate-100 via-slate-200 to-indigo-300 bg-clip-text text-transparent">
            ResearchForge
          </span>
        </NuxtLink>

        <nav class="flex items-center gap-6 text-sm font-medium">
          <NuxtLink to="/" class="text-slate-400 hover:text-slate-200 transition-colors" active-class="text-indigo-400 font-semibold">
            홈
          </NuxtLink>
          <NuxtLink to="/research/new" class="text-slate-400 hover:text-slate-200 transition-colors" active-class="text-indigo-400 font-semibold">
            새 리서치
          </NuxtLink>
          <NuxtLink to="/history" class="text-slate-400 hover:text-slate-200 transition-colors" active-class="text-indigo-400 font-semibold">
            히스토리
          </NuxtLink>
          <NuxtLink
            to="/research/new"
            class="hidden sm:inline-flex items-center gap-1.5 px-3.5 py-1.5 rounded-lg bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-semibold shadow-sm transition-all"
          >
            + 빠른 리서치
          </NuxtLink>
        </nav>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-1 max-w-6xl w-full mx-auto px-4 sm:px-6 py-8">
      <NuxtPage />
    </main>

    <!-- Footer -->
    <footer class="border-t border-slate-800/80 py-6 text-center text-xs text-slate-500">
      <div class="max-w-6xl mx-auto px-4 flex flex-col sm:flex-row items-center justify-between gap-2">
        <p>© 2026 ResearchForge. AI Multi-Agent Research Platform.</p>
        <p class="text-slate-600">Nuxt 3 · Go Fiber · FastAPI · CrewAI</p>
      </div>
    </footer>
  </div>
</template>
