const express = require('express');
const cors = require('cors');
const mongoose = require("mongoose");

const connectDb = require("./src/connection");
const UserModel = require("./src/User.model");

const app = express();
const PORT = 8080;

const whitelist = ['http://localhost:3000'];

app.use(cors({
  credentials: true,
  /*origin: function (origin, callback) {
    if (whitelist.indexOf(origin) !== -1) {
      callback(null, true)
    } else {
      callback(new Error('Not allowed by CORS'))
    }
  }*/
}));

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({
  extended: true
}));

app.get("/test", async (req, res) => {

  const id = mongoose.Types.ObjectId();

  const user = new UserModel({
    _id: id,
    name: "Bob Johnston",
    DOB: "1993-02-06T19:00:00.000Z",
    user_description: "I am Bob, i like programming and to poo",
    user_profile_picture: "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50",
    user_tags: ["Python", "Javascript", "React", "Flask", "Agile", "Poo"],
    searchingfor_tags: ["backend", "node", "Javascript"],
    user_likes: [ "482", "33", "324", "722"],
    user_liked_by: ["432", "722", "530", "348", "33"],
    common_likes: ["722", "33"],
    notified_common: ["33"],
    unnotified_common: ["722"]
});

  await user.save()
    .then(()=>{
      console.log('create user: ', user);
  })
  res.send('test');
});

app.listen(PORT, function() {
    console.log(`Running on http://localhost:${PORT}`);
    connectDb().then(() => {
      console.log("MongoDb connected");
    });
});