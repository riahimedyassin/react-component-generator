import { ExecutionContext, IGenerator, Type } from "@interfaces";
import { DirectoryService, FileService } from "@lib";
import { RCE, RFCE } from "@templates";
import { join } from "path";
import { ContextLoader } from "../context/Context.loader";

export class ComponentGenerator implements IGenerator {
  private readonly context: ExecutionContext;
  private readonly directoryService: DirectoryService;
  private readonly fileService: FileService;
  private readonly componentPath: string;
  constructor() {
    this.context = ContextLoader.context;
    this.directoryService = new DirectoryService();
    this.fileService = new FileService();
    this.componentPath = join(process.cwd(), this.context.dist);
  }
  private async generateFolder() {
    await this.directoryService.create(this.componentPath, this.context.name);
  }
  private async generateFile(extension: string, data?: string) {
    this.fileService.create(
      this.componentPath,
      `${this.context.name}.${extension}`,
      data
    );
  }
  private getTemplate() {
    return (this.context.type == Type.CLASS ? RCE : RFCE).replaceAll(
      "REPLACE_NAME",
      this.context.name
    );
  }
  public async generate(): Promise<void> {
    await this.generateFolder();
    this.generateFile(this.context.style);
    this.generateFile(this.context.template, this.getTemplate());
    if (this.context.test) {
      this.generateFile(this.context.test_extension);
    }
  }
}
