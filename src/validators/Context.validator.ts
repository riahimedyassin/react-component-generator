import { ValidatorInterface } from "@interfaces";

export class ContextValidator implements ValidatorInterface {
  validate(): Promise<void> | void {}
}
