import { ConfigGenerator, ContextLoader } from "@generators";

export class Application {
  private readonly contextLoader: ContextLoader;
  private readonly configGenerator: ConfigGenerator;
  constructor() {
    this.contextLoader = new ContextLoader();
    this.configGenerator = ConfigGenerator.getInstance();
  }

  async run() {
    await this.configGenerator.generate();
    await this.contextLoader.load();
  }
}
