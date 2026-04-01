<script setup lang="ts">
import { onMounted, watch } from "vue";
import { useRoute } from "vitepress";

const storageKey = "plugin-kit-ai-docs-locale";
const route = useRoute();

function localeFromPath(path: string): string | null {
  const normalized = path === "/" ? "/" : path.replace(/\/+$/, "");
  if (normalized === "/en" || normalized.startsWith("/en/")) {
    return "en";
  }
  if (normalized === "/ru" || normalized.startsWith("/ru/")) {
    return "ru";
  }
  return null;
}

function persistLocale(path: string) {
  if (typeof window === "undefined") {
    return;
  }
  const locale = localeFromPath(path);
  if (!locale) {
    return;
  }
  try {
    window.localStorage.setItem(storageKey, locale);
  } catch {
    // localStorage is optional enhancement only.
  }
}

onMounted(() => {
  persistLocale(route.path);
  watch(
    () => route.path,
    (path) => persistLocale(path)
  );
});
</script>

<template>
  <span hidden aria-hidden="true"></span>
</template>
