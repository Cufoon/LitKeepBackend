const path = require("node:path");

const projectRootPath = path.resolve(__dirname, "..");
const projectBuildPath = path.resolve(projectRootPath, "./build");
const execFilePath = path.resolve(projectBuildPath, "./litkeep");

module.exports = {
  apps: [
    {
      name: "litkeep",
      script: execFilePath,
      cwd: projectRootPath,
      args: "-c prod.yaml",
      log_date_format: "YYYY-MM-DD HH:mm:ss Z",
      watch: [projectBuildPath],
    },
  ],
};
