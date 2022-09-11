<template>
    <nav class="bg-white dark:bg-black" id="navbar">
        <div class="w-full flex items-center justify-center bg-suse-dark-green dark:bg-suse-dark-green-dark">
            <ul class="flex dark:text-white ">
                <li class="py-6 mx-10" id="navbar-item-1">
                    <div class="transition ease-in-out hover:-translate-y-1 hover:scale-110 duration-300">
                        <div
                            class="rounded-[12px] bg-suse-light-green text-suse-text px-10 py-2 text-4xl dark:bg-suse-light-green-dark">
                            <a href="Home">Home</a>
                        </div>
                    </div>
                </li>
                <li class="py-6 mx-10" id="navbar-item-2">
                    <div class="transition ease-in-out hover:-translate-y-1 hover:scale-110 duration-300">
                        <div
                            class="rounded-[12px] bg-suse-light-green text-suse-text px-10 py-2 text-4xl dark:bg-suse-light-green-dark">
                            <a href="AboutMe">About me</a>
                        </div>
                    </div>
                </li>
                <li class="py-6 mx-10" id="navbar-item-3">
                    <div class="transition ease-in-out hover:-translate-y-1 hover:scale-110 duration-300">
                        <div
                            class="rounded-[12px] bg-suse-light-green text-suse-text px-10 py-2 text-4xl dark:bg-suse-light-green-dark">
                            <a href="#">Contact</a>
                        </div>
                    </div>
                </li>
                <button class="fixed py-4 right-5" id="mode-button"
                    @click="setColorTheme($colorMode.preference == 'dark' ? 'light' : 'dark')">
                    <img v-if="$colorMode.value == 'dark'" src="/img/sun.svg" width="30" height="30">
                    <img v-else src="/img/moon.svg" width="30" height="30">
                </button>
            </ul>
        </div>
    </nav>
</template>

<script setup lang="ts">
type Theme = 'light' | 'dark';
const setColorTheme = (newTheme: Theme) => {
    useColorMode().preference = newTheme
}
</script>