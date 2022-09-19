<template>
    <Navbar />
    <div class="flex items-center justify-center bg-white dark:bg-black">
        <h1 class="text-black dark:text-white">Welcome to my page!</h1>
        <p> {{ test }}</p>
    </div>
</template>

<script setup>
    const { data: test } = await useFetch("https://catfact.ninja/fact", { server: true, defer: true });
</script>