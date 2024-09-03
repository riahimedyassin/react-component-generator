import { ExecutionContext, Style, Template, Type } from "@interfaces";
import { ConfigGenerator } from "../config/Config.generator";

export class ContextLoader {
  static context: ExecutionContext;
  constructor() {}
  public loadDefaultContext() {
    const config = ConfigGenerator.Config;
    ContextLoader.context = {
      dist: config.dist,
      name: "",
      test: config.test,
      style: config.style,
      template: config.template,
      type: config.type,
    };
  }
  public async load() {
    this.loadDefaultContext();
    const args = process.argv.slice(2);
    for (const value of args) {
      if (Object.values(Template).includes(value as Template)) {
        ContextLoader.context.template = value as Template;
        continue;
      }
      if (Object.values(Style).includes(value as Style)) {
        ContextLoader.context.style = value as Style;
        continue;
      }
      if (Object.values(Type).includes(value as Type)) {
        ContextLoader.context.type = value as Type;
        continue;
      }
      if (value == "test") {
        ContextLoader.context.test = true;
        continue;
      }
      const paths = value.split("/");
      ContextLoader.context.name = paths[paths.length - 1];
      if (paths.length > 1) {
        ContextLoader.context.dist = paths.slice(0, paths.length - 1).join("/");
      }
    }
  }
}
