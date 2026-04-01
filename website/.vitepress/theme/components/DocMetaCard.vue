<script setup lang="ts">
import { computed } from "vue";
import { useData } from "vitepress";

const { lang } = useData();

const props = defineProps<{
  surface: string;
  stability: string;
  maturity?: string;
  sourceRef?: string;
  sourceHref?: string;
}>();

const labels = computed(() =>
  lang.value === "ru-RU"
    ? {
        source: "Исходник",
        surface: "Поверхность",
        stability: "Стабильность",
        maturity: "Зрелость"
      }
    : {
        source: "Source",
        surface: "Surface",
        stability: "Stability",
        maturity: "Maturity"
      }
);

const prettySurface = computed(() => formatLabel(props.surface));
const prettyStability = computed(() => formatLabel(props.stability));
const prettyMaturity = computed(() => (props.maturity ? formatLabel(props.maturity) : ""));

function formatLabel(value: string) {
  return value
    .replaceAll("-", " ")
    .replaceAll("_", " ")
    .replace(/\b\w/g, (char) => char.toUpperCase());
}
</script>

<template>
  <div class="doc-meta-card">
    <span class="doc-meta-card__chip doc-meta-card__chip--surface">{{ labels.surface }}: {{ prettySurface }}</span>
    <span class="doc-meta-card__chip" :class="`doc-meta-card__chip--${stability}`">{{ labels.stability }}: {{ prettyStability }}</span>
    <span v-if="maturity" class="doc-meta-card__chip doc-meta-card__chip--maturity">{{ labels.maturity }}: {{ prettyMaturity }}</span>
    <a v-if="sourceHref" class="doc-meta-card__source" :href="sourceHref" target="_blank" rel="noreferrer">
      {{ labels.source }}
    </a>
    <code v-else-if="sourceRef" class="doc-meta-card__source">{{ sourceRef }}</code>
  </div>
</template>
