export class InvalidConfigException extends Error {
  constructor(path: string, expected: any) {
    super(
      `Invalid config parameter, please verify ${path} - Must be ${expected}`
    );
  }
}
