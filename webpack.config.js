const HtmlWebpackPlugin = require('html-webpack-plugin')

const path = require('path');

module.exports = {
  entry: './view/js/index.js',
  output: {
    filename: './js/main.js',
    path: path.resolve(__dirname, 'view/dist')
  },
  module: {
    rules: []
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: 'Nulab Products',
      template: './view/template/index.html',
    })
  ],
  watch: true,
};
