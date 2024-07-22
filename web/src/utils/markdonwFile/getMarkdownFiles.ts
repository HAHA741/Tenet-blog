const fs = require("fs");
const path = require("path");
// import * as fs from "fs";
// import * as path from "path";
const directoryPath = path.join(__dirname, "../../docs");

function getMarkdownFiles() {
  return new Promise((resolve, reject) => {
    fs.readdir(directoryPath, (err: any, files: any) => {
      if (err) {
        return reject(err);
      }
      const markdownFiles = files.filter((file: any) => file.endsWith(".md"));
      const fileContents = markdownFiles.map((file: any) => {
        const filePath = path.join(directoryPath, file);
        const content = fs.readFileSync(filePath, "utf8");
        return {
          fileName: file,
          content: content,
        };
      });
      resolve(fileContents);
    });
  });
}
export default getMarkdownFiles;
