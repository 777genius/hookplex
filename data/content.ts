import en from "~/content/en.json";
import ru from "~/content/ru.json";
import type { LandingContent, LocalizedContent } from "~/types/content";
import type { LocaleCode } from "~/data/i18n";

export const contentByLocale: LocalizedContent = {
  en,
  ru
};

export const getContent = (locale: LocaleCode): LandingContent => {
  return contentByLocale[locale] ?? contentByLocale.en;
};
