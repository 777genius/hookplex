<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { withBase } from "vitepress";

type LocaleEntry = {
  code: string;
  title: string;
  description: string;
  href: string;
};

const storageKey = "plugin-kit-ai-docs-locale";

const locales: LocaleEntry[] = [
  {
    code: "EN",
    title: "English",
    description: "Public documentation for the global plugin-kit-ai community.",
    href: withBase("/en/")
  },
  {
    code: "RU",
    title: "Русский",
    description: "Публичная документация для русскоязычных пользователей plugin-kit-ai.",
    href: withBase("/ru/")
  }
];

const preferredCode = ref("");
const manualMode = ref(false);
const preferredLocale = computed(() => locales.find((locale) => locale.code.toLowerCase() === preferredCode.value) || null);

onMounted(() => {
  try {
    preferredCode.value = window.localStorage.getItem(storageKey) || "";
  } catch {
    preferredCode.value = "";
  }

  try {
    const query = new URLSearchParams(window.location.search);
    manualMode.value = ["1", "true", "manual"].includes((query.get("gateway") || "").toLowerCase());
  } catch {
    manualMode.value = false;
  }

  if (!manualMode.value) {
    redirectToPreferredLocale();
  }
});

function rememberLocale(code: string) {
  try {
    window.localStorage.setItem(storageKey, code.toLowerCase());
  } catch {
    // localStorage is optional enhancement only.
  }
}

function detectBrowserLocale(): string {
  const candidates = [
    ...(Array.isArray(window.navigator.languages) ? window.navigator.languages : []),
    window.navigator.language || ""
  ];
  return candidates.some((candidate) => /^ru\b/i.test(String(candidate))) ? "ru" : "en";
}

function redirectToPreferredLocale() {
  const targetCode = preferredCode.value || detectBrowserLocale();
  const targetLocale = locales.find((locale) => locale.code.toLowerCase() === targetCode.toLowerCase());
  if (!targetLocale) {
    return;
  }
  rememberLocale(targetLocale.code);
  window.location.replace(targetLocale.href);
}
</script>

<template>
  <div class="language-gateway">
    <div class="language-gateway__intro">
      <p class="language-gateway__eyebrow">plugin-kit-ai docs</p>
      <h1>Choose your language</h1>
      <p>
        This gateway stays minimal on purpose. Pick a locale to enter the public documentation.
      </p>
      <p v-if="manualMode" class="language-gateway__hint">
        Automatic locale detection is paused on this manual gateway view.
      </p>
      <div v-if="preferredLocale" class="language-gateway__preferred">
        <span>Saved locale: {{ preferredLocale.title }}</span>
        <a :href="preferredLocale.href">Open preferred locale</a>
      </div>
    </div>
    <div class="language-gateway__grid">
      <a
        v-for="locale in locales"
        :key="locale.code"
        :href="locale.href"
        class="language-gateway__card"
        @click="rememberLocale(locale.code)"
      >
        <strong>{{ locale.code }}</strong>
        <span>{{ locale.title }}</span>
        <small>{{ locale.description }}</small>
      </a>
    </div>
  </div>
</template>
