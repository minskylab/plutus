// const production = !process.env.ROLLUP_WATCH;
// const purgecss = require("@fullhuman/postcss-purgecss");

module.exports = {
  plugins: [
    require("postcss-import")(),
    require("tailwindcss"),
    require("autoprefixer")
    // Only purge css on production
  ]
};
