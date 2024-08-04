// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss',"@nuxt/icon", 'nuxt-swiper'],
  runtimeConfig: {
    public: {
      api: 'http://localhost:8080/v1',
      assets: 'http://localhost:8080'
    }
  },
  compatibilityDate: "2024-07-06",
})