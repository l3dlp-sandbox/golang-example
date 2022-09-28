const path = require('path');
const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');
const Dotenv = require('dotenv-webpack');

module.exports = merge(common, {
  output: {
    filename: '[name].[hash].js',
    publicPath: '/',
  },
  devServer: {
    historyApiFallback: true,
    disableHostCheck: true,
    port: 8000,
  },
  plugins: [new Dotenv({path:'./.env.dev'})]
});
