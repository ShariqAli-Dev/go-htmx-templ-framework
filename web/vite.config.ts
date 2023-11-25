import { defineConfig } from "vite";
import fs from "fs";
import path from "path";

function getRouteEntries() {
  const routesDirectory = path.resolve(__dirname, "src/routes");
  const entryPoints: { [name: string]: string } = {};
  const routeFiles = fs.readdirSync(routesDirectory);

  routeFiles.forEach((fileName) => {
    const fileAttributes = path.parse(fileName);
    entryPoints[fileAttributes.name] = `src/routes/${fileAttributes.base}`;
  });

  return entryPoints;
}

export default defineConfig({
  build: {
    rollupOptions: {
      input: {
        index: "src/index.ts",
        ...getRouteEntries(),
      },
    },
  },
});
