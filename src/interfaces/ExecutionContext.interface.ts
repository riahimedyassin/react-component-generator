import { Config } from "./Config.interface";

export interface ExecutionContext extends Config {
  name: string;
  test_extension : string
}
