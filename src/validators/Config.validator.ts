import { ValidatorInterface } from "@interfaces";

export class ConfigValidator implements ValidatorInterface {
  validate(): Promise<void> | void {}
}
