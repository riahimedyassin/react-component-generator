export class NullValueException extends Error {
  constructor(element: string) {
    super(`${element} is missing `);
  }
}
