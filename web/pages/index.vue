<script setup>
const runtimeConfig = useRuntimeConfig()
const { data, error } = await useFetch('/banner', {
    baseURL: runtimeConfig.public.api,
})
</script>
<template>
    <SliderComponent :images="data" />
</template>