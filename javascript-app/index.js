const express = require("express");
const _ = require("lodash");

const app = express();
const port = 3000;

app.get("/", (req, res) => {
  res.send("Hello, World! (JavaScript)");
});

app.get("/greeting", (req, res) => {
  const name = _.capitalize(req.query.name || "world");
  res.send(`Hello, ${name}!`);
});

app.listen(port, () => {
  console.log(`javascript-app listening on port ${port}`);
});
