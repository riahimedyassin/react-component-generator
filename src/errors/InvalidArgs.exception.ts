export class InvalidArgsException extends Error {
  constructor(hint?: string) {
    const message = `Invalid args - Please verfiy the args required to run this command, use --help \n ${hint}`;
    super(message);
  }
}
