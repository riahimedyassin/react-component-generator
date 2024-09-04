import { IGenerator, Config } from "@interfaces";
import { FileService } from "@lib";
import { DefaultConfigTemplate } from "@templates";

export class ConfigGenerator implements IGenerator {
  public static instance: ConfigGenerator;
  private readonly fileService: FileService;
  public static Config: Config;
  private readonly configFileName = "config.json";
  private constructor() {
    this.fileService = new FileService();
  }
  public static getInstance(): ConfigGenerator {
    if (!ConfigGenerator.instance) {
      this.instance = new ConfigGenerator();
    }
    return ConfigGenerator.instance;
  }
  private async generateDefaultConfig() {
    await this.fileService.create(
      process.cwd(),
      this.configFileName,
      DefaultConfigTemplate
    );
  }
  public async generate(): Promise<void> {
    if (!this.fileService.exists(process.cwd(), this.configFileName)) {
      await this.generateDefaultConfig();
    }
    ConfigGenerator.Config = JSON.parse(
      await this.fileService.read(process.cwd(), this.configFileName)
    );
  }
}
