const path = require('path');
const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');
const Dotenv = require('dotenv-webpack');

module.exports = merge(common, {
  output: {
    filename: '[name].[hash].js',
    // Setting to the static folder outside of client so that the go server can use it to load from the server
    path: path.resolve(__dirname, '../apps/web'),
    publicPath: './',
  },
  plugins: [new Dotenv({path:'./.env.prod'})]
});
