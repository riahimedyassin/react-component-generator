import { InvalidConfigException, NullValueException } from "@errors";
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
    const { name } = this.context;
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
