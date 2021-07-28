
const {MongoClient} = require('mongodb');
var assert = require('assert');

const user = {
    "userID": {
        "$numberLong": "533"
    },
    "name": "Bob Johnston",
    "DOB": {
        "$date": "1993-02-06T19:00:00.000Z"
    },
    "user_description": "I am Bob, i like programming and to poo",
    "user_tags": ["Python", "Javascript", "React", "Flask", "Agile", "Poo"],
    "searchingfor_tags": ["backend", "node", "Javascript"],
    "user_likes": [{
        "$numberLong": "342"
    }, "482", "33", "324", "722"],
    "user_liked_by": ["432", "722", "530", "348", "33"],
    "common_likes": ["722", "33"],
    "notified_common": ["33"],
    "unnotified_common": ["722"]
}

async function main(){
   
    const uri = "mongodb+srv://Hussar:Hussar1@hussarcluster.cokdm.mongodb.net/test";
    const client = new MongoClient(uri);
 
    try {
        await client.connect(function(err, db){
            if(err) throw err;
            var dbo = db.db("Users");
            dbo.collection("Users_collection").insertOne(user,function(err,res){
                if(err) throw err;
                console.log("insert success");
                console.log(res);
                
            });
            console.log(dbo.collection("Users_collection").find().toArray(function(err, docs){
                console.log(docs);
            }));
        });

        await  listDatabases(client);


    } catch (e) {
        console.error(e);
    } finally {
        await client.close();
    }
    
}

main().catch(console.error);

async function listDatabases(client){
    databasesList = await client.db().admin().listDatabases();
 
    console.log("Databases:");
    databasesList.databases.forEach(db => console.log(` - ${db.name}`));
};