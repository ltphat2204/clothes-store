<script setup>
const { keyword } = defineProps(['keyword']);
</script>

<template>
    <div class="bg-gray-200 px-4 py-2 rounded-full text-sm flex items-center gap-1">{{ keyword }} <HotKeywordIcon/></div>
</template>