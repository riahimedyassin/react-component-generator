import { INIT_FILES, MAIN_FILES } from "@constants";
import { SupportedDB, SupportedORM } from "@enums";
import { FileNotFound, Unsupported } from "@errors";
import { ConfigInterface } from "@interfaces";
import { FileService } from ".";

export class ConfigService {
  public static config: ConfigInterface;
  private fileService: FileService;
  public static packagePath: string;
  public setPackagePath(path: string) {
    ConfigService.packagePath = path;
  }
  constructor() {
    this.fileService = FileService.getInstance();
  }
  public async load(): Promise<boolean> {
    if (!this.fileService.exists(process.cwd(), "config.json")) {
      const result = await this.init();
      if (!result) return;
    }
    const config = await this.fileService.read(process.cwd(), "config.json");
    ConfigService.config = JSON.parse(config);
    this.verifyConfig();
  }

  // Need to be moved to the validators
  // Parsing needed
  private verifyInitFiles() {
    for (const file of INIT_FILES) {
      if (
        !this.fileService.exists(
          `${ConfigService.packagePath}/src/templates/base`,
          file
        )
      ) {
        throw new FileNotFound(file);
      }
    }
    for (const file of MAIN_FILES) {
      if (
        !this.fileService.exists(
          `${ConfigService.packagePath}/src/templates/main`,
          file.dist
        )
      ) {
        throw new FileNotFound(file.dist);
      }
    }
  }

  // To validators
  private verfiyDBConfig() {
    const { ORM, type } = ConfigService.getDatabaseConfig();
    if (!Object.values(SupportedORM).includes(ORM))
      throw new Unsupported(`ORM ${ORM} is not supported yet`);
    if (!Object.values(SupportedDB).includes(type))
      throw new Unsupported(`Database ${type} is not supported yet`);
  }
  // To validators
  private verifyEntityConfig() {}

  private verifyConfig() {
    this.verfiyDBConfig();
    this.verifyInitFiles();
    this.verifyEntityConfig();
  }

  public async init() {
    return this.fileService.copy(
      `${ConfigService.packagePath}/src/templates`,
      process.cwd(),
      "config.json"
    );
  }

  public static getMainFolders() {
    return ConfigService.config.folders;
  }
  public static getAuthConfig() {
    return ConfigService.config.auth;
  }

  public static getDatabaseConfig() {
    return ConfigService.config.database;
  }
  public static getEntities() {
    return ConfigService.config.entities;
  }
}
