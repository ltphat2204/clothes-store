<script setup>
const isMobileOpen = ref(false);
const isInputFocus = ref(false);
const isSearchOpen = ref(false);

const hotKeys = computed(() => {
    return ['Áo', 'Quần', 'Váy', 'Túi', 'Giày', 'Phụ kiện', 'Sale', 'New in'];
});
</script>
<template>
    <header class="px-5 py-3 flex justify-between text-2xl">
        <div class="flex items-center">
            <Icon name="quill:hamburger-sidebar" @click="isMobileOpen = true" class="md:hidden"/>
            <NuxtLink to="/" class="h-8 w-auto">
                <img src="/brand.webp" alt="Brand Logo" class="h-full w-auto"/>
            </NuxtLink>
        </div>
        <nav class="bg-white fixed top-0 left-0 w-screen h-screen md:h-fit md:w-fit md:relative z-10 md:z-0 text-lg md:text-base font-semibold md:font-normal px-16 py-24 md:p-0 transition-transform duration-500 ease-in-out" :class="isMobileOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0'">
            <Icon name="akar-icons:cross" class="absolute top-4 right-8 md:hidden" @click="isMobileOpen = false"/>
            <ul class="flex flex-col md:flex-row gap-4 md:items-center items-start">
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
        <div class="flex items-center gap-3 md:z-10 relative">
            <Icon name="material-symbols:shopping-bag-outline"/>
            <Icon name="material-symbols:account-circle-full"/>
            <Icon name="material-symbols:search" class="hidden md:block text-3xl cursor-pointer" @click="isSearchOpen = true"/>
        </div>
    </header>
    <div class="md:fixed relative bg-black/50 w-screen md:h-screen h-fit flex justify-end top-0 transition-all duration-500 ease-in-out" :class="isSearchOpen ? 'opacity-100 z-20' : 'md:opacity-0 z-0'">
        <div class="relative w-full max-w-[500px] bg-white h-fit md:h-screen md:py-16 md:px-8 p-0 transition-all duration-500 ease-in-out" :class="isSearchOpen ? 'translate-x-0' : 'translate-x-0 md:translate-x-full'">
            <div class="mb-8 items-center gap-2 cursor-pointer hidden md:flex" @click="isSearchOpen = false"><Icon name="material-symbols:arrow-left-alt" class="text-2xl"/>Thoát</div>
            <div class="border rounded-lg px-3 py-1 mx-5 flex items-center gap-2" :class="isInputFocus ? 'border-black' : 'border-gray-300'">
                <Icon name="material-symbols:search" class="text-2xl"/>
                <input type="text" placeholder="Tìm kiếm sản phẩm" class="outline-none text-sm w-home-input" @focus="isInputFocus = true" @blur="isInputFocus = false"/>
                <Icon name="akar-icons:cross" class="text-2xl" v-if="isInputFocus" @click="isInputFocus = false"/>
            </div>
            <div class="py-2 px-5 absolute md:relative top-full left-0 md:top-0 md:left-0 w-full" v-if="isInputFocus || isSearchOpen">
                <div class="my-2">
                    Từ khóa hot
                </div>
                <div class="flex flex-wrap gap-2">
                    <HotKeyword :keyword="hotKey" v-for="hotKey in hotKeys"/>
                </div>
            </div>
        </div>
    </div>
</template>