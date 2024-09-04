export interface ValidatorInterface {
  validate(): Promise<void> | void;
}
