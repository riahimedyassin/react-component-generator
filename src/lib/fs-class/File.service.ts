import * as fsSync from "fs";
import * as fs from "fs/promises";

/**
 * @todo Better error handling
 * @todo Parse params and check for security threats
 */
export class FileService {
  public static instance: FileService;
  public static getInstance(): FileService {
    if (!FileService.instance) {
      FileService.instance = new FileService();
    }
    return FileService.instance;
  }
  async copy(src: string, dist: string, name: string) {
    return this.read(src, name)
      .then(async (data) => {
        return this.create(dist, name, data)
          .then((response) => response)
          .catch((_) => {
            return false;
          });
      })
      .catch((_) => {
        return false;
      });
  }
  exists(path: string, name: string) {
    return fsSync.existsSync(`${path}/${name}`);
  }

  async create(path: string, name: string, data: string = "") {
    return fs
      .writeFile(`${path}/${name}`, data)
      .then((_) => true)
      .catch((_) => false);
  }
  async read(path: string, name: string): Promise<string> {
    return fs.readFile(`${path}/${name}`, { encoding: "utf-8", flag: "r" });
  }
  async overwrite(path: string, name: string, value: string) {
    await fs.writeFile(`${path}/${name}`, value);
  }
}
