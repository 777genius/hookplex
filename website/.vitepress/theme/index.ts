import DefaultTheme from "vitepress/theme";
import type { Theme } from "vitepress";
import LanguageGateway from "./components/LanguageGateway.vue";
import DocMetaCard from "./components/DocMetaCard.vue";
import NotFoundLinks from "./components/NotFoundLinks.vue";
import "./styles/custom.css";

const theme: Theme = {
  ...DefaultTheme,
  enhanceApp({ app }) {
    app.component("LanguageGateway", LanguageGateway);
    app.component("DocMetaCard", DocMetaCard);
    app.component("NotFoundLinks", NotFoundLinks);
  }
};

export default theme;
