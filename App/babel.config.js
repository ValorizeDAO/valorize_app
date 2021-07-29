
module.exports = {
  plugins: [
    "babel-plugin-transform-import-meta",
    function () {
      return {
        visitor: {
          MetaProperty(path) {
            path.replaceWithSourceString("process");
          },
        },
      }
    },
  ],
}
