const mongoose = require("mongoose");

const DB_NAME = 'Users'
const connection = `mongodb+srv://Hussar:Hussar1@hussarcluster.cokdm.mongodb.net/${DB_NAME}`;

const connectDb = () => {
    return mongoose.connect(connection, { useNewUrlParser: true,  useUnifiedTopology: true});
};

module.exports = connectDb;