<script setup lang="ts">
import { computed } from "vue";
import { useData } from "vitepress";
import DefaultTheme from "vitepress/theme";
import { useMermaidZoom } from "vitepress-mermaid-zoom";
import { useCodeblockCollapse } from "vitepress-codeblock-collapse";
import LocalePreferenceSync from "./LocalePreferenceSync.vue";
import "vitepress-mermaid-zoom/style.css";
import "vitepress-codeblock-collapse/style.css";

const { page } = useData();
const pagePath = computed(() => page.value.relativePath);

useMermaidZoom(pagePath, {
  selector: ".mermaid-diagram__surface"
});

useCodeblockCollapse(pagePath);
</script>

<template>
  <DefaultTheme.Layout>
    <template #layout-top>
      <LocalePreferenceSync />
    </template>
  </DefaultTheme.Layout>
</template>
