const express = require("express");
const bodyParser = require("body-parser");
const app = express();

app.use(bodyParser.json());

let users = require("./db.json").users;

app.post("/users", (req, res) => {
    const { name } = req.body;
    const user = { id: users.length + 1, name };
    users.push(user);
    res.json(user);
});

app.listen(3000, () => console.log("Server started"));