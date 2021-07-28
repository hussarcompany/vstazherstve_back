const mongoose = require("mongoose");

const options = {
  collection: 'users'
};

const objectSchema = new mongoose.Schema({
  name: {
    type: String,
    default: 'unnamed'
  },
  DOB: {
    type: Date
  },
  user_description: {
    type: String,
    default: 'Нет'
  },
  user_tags: {
    type: [String],
    default: ['Нет']
  },
  searchingfor_tags: {
    type: [String],
    default: ['Нет']
  },
  user_likes: {
    type: [String],
    default: ['Нет']
  },
  user_liked_by: {
    type: [String],
    default: ['Нет']
  },
  common_likes: {
    type: [String],
    default: ['Нет']
  },
  notified_common: {
    type: [String],
    default: ['Нет']
  },
  unnotified_common: {
    type: [String],
    default: ['Нет']
  },
}, options);

const UserModel = mongoose.model("Users", objectSchema);

module.exports = UserModel;