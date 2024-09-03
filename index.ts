import { Application } from "src/app";

export async function bootstrap() {
  const app = new Application();
  await app.run();
}

bootstrap();
