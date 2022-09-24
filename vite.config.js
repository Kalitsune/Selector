import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      '/auth': 'http://127.0.0.1:8080',
      '/api': 'http://127.0.0.1:8080',
    }
  },
  plugins: [vue()]
})
