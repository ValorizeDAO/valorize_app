module.exports = {
  preset: "@vue/cli-plugin-unit-jest/presets/no-babel",
  moduleFileExtensions: ["js", "ts", "json", "vue"],
  transform: {
    "^.+\\.ts$": "ts-jest",
    "^.+\\.vue$": "vue-jest",
  },
  testURL: "http://localhost/",
  globals: {
    "vue-jest": {
      babelConfig: true,
    },
  },
};
