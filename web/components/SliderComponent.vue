<script setup>
import { Navigation } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/navigation';

const runtimeConfig = useRuntimeConfig();
const props = defineProps({
    images: Object,
});

const data = props.images?.data;
useSwiper;
</script>

<template>
    <div class="w-full max-w-[1320px] mx-auto">
        <Swiper
            :modules="[SwiperAutoplay, SwiperEffectCreative, Navigation]"
            :slides-per-view="1"
            :effect="'creative'"
            :loop="true"
            :autoplay="{
                delay: 2000,
                disableOnInteraction: true,
            }"
            :creative-effect="{
                prev: {
                    shadow: false,
                    translate: ['-100%', 0, -1],
                },
                next: {
                    translate: ['100%', 0, 0],
                },
            }"
        >
            <SwiperSlide v-for="(banner, index) in data" :key="index">
                <NuxtLink :to="banner.redirect_url">
                    <img class="hidden md:block" :src="`${runtimeConfig.public.assets}${banner.image}`" alt="Banner" />
                    <img class="block md:hidden" :src="`${runtimeConfig.public.assets}${banner.image_small}`" alt="Mobile banner" />
                </NuxtLink>

            </SwiperSlide>
        </Swiper>
    </div>
</template>