import { ConfigGenerator, ContextLoader } from "@generators";
import { ContextValidator } from "./validators/Context.validator";
import { ConfigValidator } from "./validators/Config.validator";
import { ComponentGenerator } from "./generators/components/Components.generator";

export class Application {
  private readonly contextLoader: ContextLoader;
  private readonly configGenerator: ConfigGenerator;
  private readonly contextValidator: ContextValidator;
  private readonly configValidator: ConfigValidator;
  private componentGenerator!: ComponentGenerator;
  constructor() {
    this.contextLoader = new ContextLoader();
    this.configGenerator = ConfigGenerator.getInstance();
    this.configValidator = new ConfigValidator();
    this.contextValidator = new ContextValidator();
  }
  private getComponentGenerator() {
    this.componentGenerator = new ComponentGenerator();
    return this.componentGenerator;
  }

  async run() {
    await this.configGenerator.generate();
    await this.contextLoader.load();
    this.configValidator.setConfig(ConfigGenerator.Config).validate();
    this.contextValidator.setContext(ContextLoader.context).validate();
    this.getComponentGenerator().generate();
  }
}
