import { ConfigGenerator, ContextLoader } from "@generators";
import { ContextValidator } from "./validators/Context.validator";
import { ConfigValidator } from "./validators/Config.validator";
import { ComponentGenerator } from "./generators/components/Components.generator";
import { InvalidConfigException, NullValueException } from "@errors";

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
    try {
      await this.configGenerator.generate();
      await this.contextLoader.load();
      if (ContextLoader.context.name == "--init") {
        console.log("Config file generated successfully");
        return;
      }
      this.configValidator.setConfig(ConfigGenerator.Config).validate();
      this.contextValidator.setContext(ContextLoader.context).validate();
      this.getComponentGenerator().generate();
    } catch (error) {
      if (error instanceof InvalidConfigException) {
        console.log(`Validation Failed : ${error.message}`);
        return;
      }
      if (error instanceof NullValueException) {
        console.log(`Null value exception : ${error.message}`);
        return;
      }
      console.log(error);
    }
  }
}
