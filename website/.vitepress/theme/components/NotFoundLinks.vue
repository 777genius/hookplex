<script setup lang="ts">
import { withBase } from "vitepress";

const links = [
  {
    title: "Language Gateway",
    description: "Return to the entry page and choose a locale again.",
    href: withBase("/")
  },
  {
    title: "English Docs",
    description: "Jump into the English public docs root.",
    href: withBase("/en/")
  },
  {
    title: "Русская документация",
    description: "Откройте русскоязычный корень публичной документации.",
    href: withBase("/ru/")
  }
];
</script>

<template>
  <div class="docs-grid">
    <a v-for="link in links" :key="link.href" class="docs-card" :href="link.href">
      <h2>{{ link.title }}</h2>
      <p>{{ link.description }}</p>
    </a>
  </div>
</template>
