const path = require("path");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const { AlwaysDepth } = require("three");

let imageSizeLimit = 9007199254740991; // Number.MAX_SAFE_INTEGER
let sourceDir = path.resolve(__dirname, "src");
let buildDir = path.resolve(__dirname, "../dist");

module.exports = {
  entry: {
    index: path.resolve(sourceDir, "main.ts"),
  },
  output: {
    path: buildDir,
    filename: "main.js",
  },
  optimization: {
    splitChunks: false,
  },
  devServer: {
    static: {
      directory: path.join(__dirname, 'src'),
    },
    compress: true,
    port: 9000,
  },
  performance: {
    hints: false,
  },
  mode: "production",
  module: {
    rules: [
      {
        test: /\.(png|gif|jpg|woff2?|eot|ttf|otf|svg)(\?.*)?$/i,
        use: [
          {
            loader: "url-loader",
            options: {
              limit: imageSizeLimit,
            },
          },
        ],
      },
      {
        test: /\.ts$/,
        use: "ts-loader"
      }
    ],
  },
  resolve: {
    extensions: [".ts", ".js", ".json"]
  },
  plugins: [
    new CopyWebpackPlugin({
      patterns: [
        {
          from: path.resolve(sourceDir, "main.css"),
          to: path.resolve(buildDir, "main.css"),
        },
        {
          from: path.resolve(sourceDir, "index.html"),
          to: path.resolve(buildDir, "index.html"),
        },
        {
          from: path.resolve(sourceDir, "404.html"),
          to: path.resolve(buildDir, "404.html")
        }
      ],
    }),
  ],
};
