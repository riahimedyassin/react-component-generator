import { Application } from "src/app";

export function bootstrap() {
  const app = new Application(__dirname);
  app.run();
}

bootstrap();
