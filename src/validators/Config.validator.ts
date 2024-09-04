import { InvalidConfigException, NullValueException } from "@errors";
import { Config, Style, Template, Type, ValidatorInterface } from "@interfaces";

export class ConfigValidator implements ValidatorInterface {
  constructor() {}
  private config!: Config;
  public setConfig(config: Config) {
    this.config = config;
    return this;
  }
  public validate(): void {
    const VALID_KEYS = new Set(["template", "style", "type", "dist", "test"]);
    const keys = Object.keys(this.config);
    for (const key of keys) {
      if (!VALID_KEYS.has(key))
        throw new InvalidConfigException("keys", VALID_KEYS);
    }
    if (!this.config) {
      throw new NullValueException("config");
    }
    if (!Object.values(Template).includes(this.config.template)) {
      throw new InvalidConfigException("template", Template);
    }
    if (!Object.values(Style).includes(this.config.style)) {
      throw new InvalidConfigException("style", Style);
    }
    if (!Object.values(Type).includes(this.config.type)) {
      throw new InvalidConfigException("type", Type);
    }
    if (!(typeof this.config.test === "boolean")) {
      throw new InvalidConfigException("test", "boolean");
    }
    const { dist } = this.config;
    if (!(dist.trim().length == dist.length && dist.toLowerCase() == dist)) {
      throw new InvalidConfigException(
        "dist",
        "All lower case & No white spaces"
      );
    }
  }
}
