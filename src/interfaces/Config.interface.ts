export enum Style {
  SCSS = "scss",
  CSS = "css",
  LESS = "less",
  SASS = "sass",
}

export enum Template {
  JAVASCRIPT = "jsx",
  TYPESCRIPT = "tsx",
}
export enum Type {
  FUNCTION = "rfce",
  CLASS = "rce",
}

export interface Config {
  style: Style;
  template: Template;
  type: Type;
  dist: string;
  test: boolean;
}
