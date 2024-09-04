import { ExecutionContext, IGenerator, Type } from "@interfaces";
import { DirectoryService, FileService } from "@lib";
import { RCE, RFCE } from "@templates";
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
    this.componentPath = `${process.cwd()}\\${this.context.dist}`;
  }
  private async generateFolder() {
    await this.directoryService.create(this.componentPath, this.context.name);
  }
  private async generateParentDirectories() {
    const directories = this.context.dist.split("\\");
    let stack = "";
    for (const directory of directories) {
      const path = `${process.cwd()}\\${stack != "" ? stack + "\\" : ""}`;
      if (!(await this.directoryService.exists(path, directory))) {
        await this.directoryService.create(path, directory);
      }
      stack += directory;
    }
  }
  private async generateFile(extension: string, data?: string) {
    await this.fileService.create(
      `${this.componentPath}\\${this.context.name}`,
      `${this.context.name}.${extension}`,
      data
    );
    console.log(
      `Generated ${extension} File at ${this.componentPath}/${this.context.name}`
    );
  }
  private getTemplate() {
    return (this.context.type == Type.CLASS ? RCE : RFCE).replaceAll(
      "REPLACE_NAME",
      this.context.name
    );
  }
  public async generate(): Promise<void> {
    await this.generateParentDirectories();
    await this.generateFolder();
    await this.generateFile(this.context.style);
    await this.generateFile(this.context.template, this.getTemplate());
    if (this.context.test) {
      await this.generateFile(this.context.test_extension);
    }
  }
}
