<script setup lang="ts">
const { content } = useLandingContent()
const { t } = useI18n()

function onGetStarted() {
  const el = document.getElementById("download")
  if (el) {
    el.scrollIntoView({ behavior: "smooth" })
  }
}
</script>

<template>
  <section id="pricing" class="pricing-section section anchor-offset">
    <v-container>
      <div class="pricing-section__header">
        <h2 class="pricing-section__title">
          {{ t("pricing.sectionTitle") }}
        </h2>
        <p class="pricing-section__subtitle">
          {{ t("pricing.sectionSubtitle") }}
        </p>
      </div>

      <div class="pricing-section__cards">
        <article
          v-for="plan in content.pricing"
          :key="plan.id"
          class="pricing-card"
          :class="{ 'pricing-card--highlighted': plan.highlighted }"
        >
          <div class="pricing-card__popular">
            {{ t("pricing.popular") }}
          </div>

          <div class="pricing-card__header">
            <h3 class="pricing-card__name">{{ plan.name }}</h3>
            <div class="pricing-card__price-wrap">
              <span class="pricing-card__price">{{ plan.price }}</span>
              <span class="pricing-card__period">/ {{ plan.period }}</span>
            </div>
            <p class="pricing-card__description">
              {{ plan.description }}
            </p>
          </div>

          <ul class="pricing-card__features">
            <li
              v-for="feature in plan.features"
              :key="feature"
              class="pricing-card__feature"
            >
              <span class="pricing-card__feature-icon" aria-hidden="true">+</span>
              <span>{{ feature }}</span>
            </li>
          </ul>

          <button
            type="button"
            class="pricing-card__btn pricing-card__btn--primary"
            @click="onGetStarted"
          >
            {{ t("pricing.getStarted") }}
          </button>
        </article>
      </div>

      <p class="pricing-section__note">
        <span class="pricing-section__note-badge">OSI</span>
        {{ t("pricing.note") }}
      </p>
    </v-container>
  </section>
</template>

<style scoped>
.pricing-section {
  position: relative;
}

.pricing-section__header {
  text-align: center;
  max-width: 640px;
  margin: 0 auto 56px;
  position: relative;
  z-index: 1;
}

.pricing-section__title {
  font-size: 2.4rem;
  font-weight: 800;
  letter-spacing: -0.03em;
  line-height: 1.15;
  margin-bottom: 16px;
  background: linear-gradient(135deg, #e0e6ff 0%, #39ff14 100%);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.pricing-section__subtitle {
  font-size: 1.1rem;
  color: #8892b0;
  line-height: 1.6;
  margin: 0;
}

.pricing-section__cards {
  max-width: 760px;
  margin: 0 auto;
}

.pricing-card {
  position: relative;
  background: rgba(10, 10, 15, 0.82);
  border: 1px solid rgba(0, 240, 255, 0.18);
  border-radius: 24px;
  padding: 36px 30px;
  display: flex;
  flex-direction: column;
  gap: 28px;
  backdrop-filter: blur(12px);
  box-shadow: 0 0 40px rgba(0, 240, 255, 0.06);
}

.pricing-card--highlighted {
  border-color: rgba(57, 255, 20, 0.28);
  background: linear-gradient(
    180deg,
    rgba(0, 240, 255, 0.08) 0%,
    rgba(57, 255, 20, 0.05) 100%
  );
}

.pricing-card__popular {
  position: absolute;
  top: -1px;
  right: 24px;
  padding: 6px 16px;
  border-radius: 0 0 12px 12px;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  background: linear-gradient(135deg, #00f0ff, #39ff14);
  color: #0a0a0f;
  font-family: "JetBrains Mono", monospace;
}

.pricing-card__header {
  display: grid;
  gap: 12px;
}

.pricing-card__name {
  font-size: 1.2rem;
  font-weight: 700;
  margin: 0;
  color: #e0e6ff;
  font-family: "JetBrains Mono", monospace;
}

.pricing-card__price-wrap {
  display: flex;
  align-items: baseline;
  gap: 4px;
}

.pricing-card__price {
  font-size: 3rem;
  font-weight: 800;
  letter-spacing: -0.03em;
  line-height: 1;
  background: linear-gradient(135deg, #00f0ff, #39ff14);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.pricing-card__period {
  font-size: 0.9rem;
  opacity: 0.58;
  font-weight: 500;
  color: #8892b0;
}

.pricing-card__description {
  font-size: 0.95rem;
  color: #a8b3d1;
  line-height: 1.6;
  margin: 0;
}

.pricing-card__features {
  margin: 0;
  padding: 0;
  list-style: none;
  display: grid;
  gap: 12px;
}

.pricing-card__feature {
  display: grid;
  grid-template-columns: 18px minmax(0, 1fr);
  gap: 12px;
  align-items: start;
  color: #dce6ff;
  line-height: 1.55;
}

.pricing-card__feature-icon {
  color: #39ff14;
  font-weight: 800;
}

.pricing-card__btn--primary {
  width: 100%;
  padding: 14px 24px;
  border-radius: 12px;
  font-size: 0.95rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.25s ease;
  background: linear-gradient(135deg, #00f0ff, #ff00ff);
  color: #0a0a0f;
  border: none;
  box-shadow: 0 4px 20px rgba(0, 240, 255, 0.3);
  font-family: "JetBrains Mono", monospace;
  letter-spacing: 0.02em;
}

.pricing-card__btn--primary:hover {
  box-shadow: 0 6px 30px rgba(0, 240, 255, 0.5);
  transform: translateY(-1px);
}

.pricing-section__note {
  text-align: center;
  margin-top: 32px;
  font-size: 0.85rem;
  color: #8892b0;
  opacity: 0.78;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  position: relative;
  z-index: 1;
}

.pricing-section__note-badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 34px;
  height: 22px;
  border-radius: 999px;
  font-size: 0.68rem;
  letter-spacing: 0.08em;
  font-weight: 700;
  background: rgba(57, 255, 20, 0.14);
  color: #39ff14;
  font-family: "JetBrains Mono", monospace;
}

.v-theme--light .pricing-section__title {
  background: linear-gradient(135deg, #1e293b 0%, #059669 100%);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.v-theme--light .pricing-section__subtitle,
.v-theme--light .pricing-section__note {
  color: #475569;
}

.v-theme--light .pricing-card {
  background: rgba(255, 255, 255, 0.84);
  border-color: rgba(14, 165, 233, 0.18);
}

.v-theme--light .pricing-card__name,
.v-theme--light .pricing-card__feature {
  color: #1e293b;
}

.v-theme--light .pricing-card__description,
.v-theme--light .pricing-card__period {
  color: #475569;
}

@media (max-width: 960px) {
  .pricing-section__title {
    font-size: 1.85rem;
  }

  .pricing-section__header {
    margin-bottom: 40px;
  }

  .pricing-section__subtitle {
    font-size: 1rem;
  }
}

@media (max-width: 600px) {
  .pricing-card {
    padding: 32px 22px 24px;
  }
}
</style>
