import { ref, onUnmounted } from 'vue'
import type { WebSocketEvent } from '~/types'

export const useWebSocket = (projectId: string) => {
  const config = useRuntimeConfig()
  const wsBase = config.public.wsBase

  const events = ref<WebSocketEvent[]>([])
  const currentStep = ref<string>('')
  const progress = ref<number>(0)
  const isConnected = ref<boolean>(false)
  const isFinished = ref<boolean>(false)
  const error = ref<string | null>(null)

  let ws: WebSocket | null = null

  const connect = () => {
    if (!projectId || typeof window === 'undefined') return

    const wsUrl = `${wsBase}/research/${projectId}`
    ws = new WebSocket(wsUrl)

    ws.onopen = () => {
      isConnected.value = true
      error.value = null
    }

    ws.onmessage = (event) => {
      try {
        const payload: WebSocketEvent = JSON.parse(event.data)
        events.value.push(payload)

        if (payload.progress !== undefined) {
          progress.value = payload.progress
        }

        if (payload.agent_name) {
          currentStep.value = payload.agent_name
        }

        if (payload.type === 'PROJECT_DONE') {
          isFinished.value = true
          progress.value = 100
        } else if (payload.type === 'ERROR') {
          error.value = payload.message
          isFinished.value = true
        }
      } catch (e) {
        console.error('Failed to parse WS message:', e)
      }
    }

    ws.onerror = (e) => {
      console.warn('WebSocket connection error:', e)
    }

    ws.onclose = () => {
      isConnected.value = false
    }
  }

  const disconnect = () => {
    if (ws) {
      ws.close()
      ws = null
    }
  }

  onUnmounted(() => {
    disconnect()
  })

  return {
    events,
    currentStep,
    progress,
    isConnected,
    isFinished,
    error,
    connect,
    disconnect
  }
}
