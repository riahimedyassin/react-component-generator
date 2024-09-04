import {
  InvalidArgsException,
  InvalidConfigException,
  NullValueException,
} from "@errors";
import { ExecutionContext, ValidatorInterface } from "@interfaces";

export class ContextValidator implements ValidatorInterface {
  private context!: ExecutionContext;
  setContext(ctx: ExecutionContext) {
    this.context = ctx;
    return this;
  }

  validate(): Promise<void> | void {
    if (!this.context) {
      throw new NullValueException("context");
    }
    this.validateName();
    if (!this.context.test_extension) {
      throw new InvalidConfigException("test_extension", [
        "spec.js",
        "spec.ts",
      ]);
    }
  }
  private validateName() {
    const { name } = this.context;
    if (!name) {
      throw new InvalidArgsException(
        `Provide the component name as an argument`
      );
    }
    if (!(name.trim().length === name.length)) {
      throw new InvalidConfigException(
        "name",
        "Component name must have no white spaces"
      );
    }
    if (name[0].toUpperCase() < "A" || name[0].toUpperCase() > "Z") {
      throw new InvalidConfigException(
        "name",
        "Component name must be valid name that starts with an alphabetic value"
      );
    }
  }
}
