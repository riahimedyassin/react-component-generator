import { IGenerator } from "@interfaces";
import { FileService } from "@lib";

export class ConfigGenerator implements IGenerator {
  public static instance: ConfigGenerator;
  public static executionPath: string;
  private readonly fileService: FileService;
  public static argv: string[];
  private constructor() {
    ConfigGenerator.executionPath = process.cwd();
    ConfigGenerator.argv = process.execArgv;
    this.fileService = new FileService();
  }
  public static getInstance(): ConfigGenerator {
    if (!ConfigGenerator.instance) return new this();
    return ConfigGenerator.instance;
  }
  private async loadConfig() {}
  public generate(): Promise<void> | void {}
}
