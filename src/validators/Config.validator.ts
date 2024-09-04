import { InvalidConfigException } from "@errors";
import { Config, Style, Template, Type, ValidatorInterface } from "@interfaces";
import { NullValueException } from "src/errors/NullValue.exception";

export class ConfigValidator implements ValidatorInterface {
  constructor() {}
  private config!: Config;
  public setConfig(config: Config) {
    this.config = config;
    return this;
  }
  public validate(): void {
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
