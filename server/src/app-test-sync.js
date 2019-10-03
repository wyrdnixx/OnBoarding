const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')
const morgan = require('morgan')


const app = express()
app.use(morgan('combined'))
app.use(bodyParser.json())
app.use(cors())


//const mysql = require('mysql');


var Sequelize = require('sequelize');


// var pool = mysql.createPool({
//     connectionLimit: 100,
//     host: process.env.DATABASE_URL,
//     user: 'root',
//     password : process.env.MYSQL_ROOT_PASSWORD,
//     database : process.env.MARIADB_DATABASE,
//     debug: false
    
// });


var sequelize = new Sequelize(process.env.MARIADB_DATABASE, 'root', process.env.MYSQL_ROOT_PASSWORD, {
    host: 'localhost',
    port: 3306,
    dialect: 'mysql'
});


app.get('/', async (req, res) => {



        var sql = " select \
Personal.id as PersonalId, \
Firmen.name as firma, \
persnr, \
nachname,\
vorname, \
titel, \
geschlecht, \
eintritt, \
austritt, \
Abteilungen.name as abteilung, \
Person_Abteilung_Status.status as abteilungStatus, \
gebdat, \
dienstart, \
berufsbezeichnung, \
Person_Processor_Status.status as processorStatus \
from Personal \
left join Firmen on Firmen.id = Personal.firmaId \
left join Abteilungen on Abteilungen.id = Personal.abteilungId \
left join Person_Abteilung_Status on Personal.id = Person_Abteilung_Status.personId \
left join Person_Processor_Status on Personal.id = Person_Processor_Status.personId"



//Checking connection status
sequelize.authenticate().complete(function (err) {
    if (err) {
       console.log('There is connection in ERROR');
    } else {
       console.log('Connection has been established successfully');
    }
   });



        // await  pool.query(sql, function (err, result) {
        // console.log("-----------   first  -------")
        //     if (err) throw err
        //     else 
        //         console.log(result[0].persnr)
        //         //res.json(result);
        //         //pool.release();
                
        // })
        
        //     await  pool.query("select 1+1", function (err, result) {
        //         console.log("-----------   second  -------")
        //         if (err) throw err
        //         else 
        //             console.log("-----------------------------------",result)
        //             //res.json(result);
        //             //pool.release();
        //     });      
        
        //     res.json(result);
        

});

  
app.listen(process.env.PORT || 8081, () => {
    console.log("App Server started");
  })