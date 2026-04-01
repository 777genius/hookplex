import { defineConfig } from "vitepress";
import { sharedConfig } from "./shared";
import { enLocaleConfig } from "./locales.en";
import { ruLocaleConfig } from "./locales.ru";

export default defineConfig({
  ...sharedConfig,
  locales: {
    root: sharedConfig.locales?.root,
    en: enLocaleConfig,
    ru: ruLocaleConfig
  }
});
