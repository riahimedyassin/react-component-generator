import * as fs from "fs/promises";

export class DirectoryService {
  constructor() {}
  public static instance: DirectoryService;
  public static getInstance() {
    if (!DirectoryService.instance) {
      DirectoryService.instance = new DirectoryService();
    }
    return DirectoryService.instance;
  }

  create(path: string, ...names: string[]): Promise<boolean>;
  create(path: string, names: string): Promise<boolean>;
  async create(path: string, ...names: string[]) {
    if (Array.isArray(names)) {
      names.forEach(async (name) => {
        await fs
          .mkdir(`${path}\\${name}`, { recursive: true })
          .catch((_) => false);
      });
      return true;
    } else {
      await fs
        .mkdir(`${path}\\${names}`, { recursive: true })
        .catch((_) => false);
      return true;
    }
  }
  async exists(path: string, name: string) {
    return fs
      .readdir(`${path}\\${name}`)
      .then(() => true)
      .catch(() => false);
  }
}
