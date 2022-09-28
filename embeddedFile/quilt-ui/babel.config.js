const { BABEL_ENV } = process.env;

const presets = ["@babel/preset-env", "@babel/preset-react"]
let plugins = []

if (BABEL_ENV && BABEL_ENV === "test") {
    plugins = ["file-loader"]
}

module.exports = { presets, plugins }