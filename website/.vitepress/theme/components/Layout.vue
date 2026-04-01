<script setup lang="ts">
import { computed } from "vue";
import { useData } from "vitepress";
import DefaultTheme from "vitepress/theme";
import { useMermaidZoom } from "vitepress-mermaid-zoom";
import { useCodeblockCollapse } from "vitepress-codeblock-collapse";
import LocalePreferenceSync from "./LocalePreferenceSync.vue";
import LocaleSwitcher from "./LocaleSwitcher.vue";
import "vitepress-mermaid-zoom/style.css";
import "vitepress-codeblock-collapse/style.css";

const { page } = useData();
const pagePath = computed(() => page.value.relativePath);

useMermaidZoom(pagePath);

useCodeblockCollapse(pagePath);
</script>

<template>
  <DefaultTheme.Layout>
    <template #layout-top>
      <LocalePreferenceSync />
    </template>
    <template #nav-bar-content-after>
      <LocaleSwitcher />
    </template>
    <template #nav-screen-content-after>
      <LocaleSwitcher variant="screen" />
    </template>
  </DefaultTheme.Layout>
</template>
