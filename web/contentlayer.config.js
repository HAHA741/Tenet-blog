import { defineDocumnetType, makeSource } from "contentlayer/source-files";
const Post = defineDocumnetType(() => ({
  name: "Post",
  filePathPattern: "**/*.md",
  filelds: {
    title: { type: "string", required: true },
    date: { type: "date", required: true },
    body: { type: "markdown", required: true },
  },
}));
export default makeSource({
  contentDirPath: "docs",
  documentTypes: [Post],
});
