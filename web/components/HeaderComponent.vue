<script setup>
const isMobileOpen = ref(false);
</script>
<template>
    <header class="px-5 py-3 flex justify-between text-2xl">
        <div class="flex items-center">
            <Icon name="quill:hamburger-sidebar" @click="isMobileOpen = true" class="md:hidden"/>
            <NuxtLink to="/" class="h-8 w-auto">
                <img src="/brand.webp" alt="Brand Logo" class="h-full w-auto"/>
            </NuxtLink>
        </div>
        <nav class="bg-white fixed top-0 left-0 w-screen h-screen md:h-fit md:w-fit md:relative z-50 text-lg md:text-base font-semibold md:font-normal px-16 py-24 md:p-0 transition-transform duration-500 ease-in-out" :class="isMobileOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0'">
            <Icon name="akar-icons:cross" class="absolute top-4 right-8 md:hidden" @click="isMobileOpen = false"/>
            <ul class="flex flex-col md:flex-row gap-4 items-center">
                <li>
                    <NuxtLink to="/" class="block">NEW IN</NuxtLink>
                </li>
                <li>
                    <NuxtLink to="/" class="block">SẢN PHẨM</NuxtLink>
                </li>
                <li>
                    <NuxtLink to="/" class="block">LOOKBOOK</NuxtLink>
                </li>
                <li>
                    <NuxtLink to="/" class="block">SỰ KIỆN</NuxtLink>
                </li>
                <li>
                    <NuxtLink to="/" class="block">BLOG</NuxtLink>
                </li>
                <li>
                    <NuxtLink to="/" class="block">CỬA HÀNG</NuxtLink>
                </li>
            </ul>
        </nav>
        <div class="flex items-center gap-3">
            <Icon name="material-symbols:shopping-bag-outline"/>
            <Icon name="material-symbols:account-circle-full"/>
        </div>
    </header>
</template>