export interface IGenerator {
  generate(): Promise<void> | void;
}
